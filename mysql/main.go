package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/schollz/progressbar/v3"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup
var db *sql.DB
var bar *progressbar.ProgressBar

func initDB() (err error) {
	dsn := "root:Kankan.123@tcp(47.98.63.198:3306)/bb?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		print(err.Error())
		return err
	}

	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		print(err.Error())
		return err
	}

	return nil
}

func listCode() []string {
	var codes []string
	rows, err := db.Query(`SELECT b.code FROM _zbrs b WHERE 1=1`)
	if err != nil {
		fmt.Printf("=== query failed, err: %s\n", err.Error())
		return nil
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var code string
		err := rows.Scan(&code)
		if err != nil {
			fmt.Printf("=== scan failed, err: %s\n", err.Error())
			return nil
		}
		codes = append(codes, code)
	}

	return codes
}

func setNum(codes []string) {
	defer wg.Done()

	fmt.Printf("=== 处理 codes: %v\n", codes)
	for _, code := range codes {
		// 查询人数
		var num int
		sqlStr := `
			WITH RECURSIVE _branch AS (
				SELECT b.*, b.code child_code, b.name child_name, b.deletedTime child_deletedTime, b.type child_type FROM branch b
				UNION ALL
				SELECT b.*, _b.child_code, _b.child_name, _b.child_deletedTime, _b.child_type FROM _branch _b INNER JOIN branch b ON _b.parent_code = b.code
			)
	
			SELECT COUNT(vmm.guid)
			FROM _branch _b
					 LEFT JOIN VIEW_MISC_MEMBER vmm ON _b.child_code = vmm.branch_code AND vmm.m_status IN (0, 1, 3)
			WHERE _b.code = ?
			  AND _b.deletedTime IS NULL
			  AND _b.child_deletedTime IS NULL
		`
		err := db.QueryRow(sqlStr, code).Scan(&num)
		if err != nil {
			fmt.Printf("=== get num failed, err: %s\n", err.Error())
		}

		// 更新
		sql := "UPDATE _zbrs SET num = ? WHERE code = ?"
		_, err = db.Exec(sql, num, code)
		if err != nil {
			fmt.Printf("=== 更新失败, err: %s\n", err.Error())
			return
		}
		//fmt.Printf("=== 更新成功, code: %s, num: %d\n", code, num)

		// update progressbar
		bar.Add(1)
	}
}

var GO_NUM = 4

func main() {
	// 初始化数据库
	initDB()

	// 获取所有支部code
	codes := listCode()
	//codes := []string{"1111", "2222", "3333", "4444", "5555", "6666", "7777"}

	// 初始化 progressbar
	bar = progressbar.NewOptions(len(codes),
		progressbar.OptionSetDescription("Done"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(50),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionShowIts(),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionSetRenderBlankState(true),
	)

	// 设置支部人数
	for i := 0; i < GO_NUM; i++ {
		wg.Add(1)
		go setNum(codes[i*len(codes)/GO_NUM : (i+1)*len(codes)/GO_NUM])
	}

	wg.Wait()
	fmt.Printf("=== main 结束")
}
