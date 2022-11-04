package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

// REF: https://zhuanlan.zhihu.com/p/442534091
var wg sync.WaitGroup
var db *sql.DB

func initDB() (err error) {
	dsn := "test_check:Renkankan@2020@tcp(10.119.0.254:3306)/zhdj?charset=utf8mb4&parseTime=True"
	//dsn := "zhdj:95CED76469FEA636@tcp(192.168.2.129:3306)/zhdj?charset=utf8mb4&parseTime=True"
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
		fmt.Printf("=== 更新成功, code: %s, num: %d\n", code, num)
	}
}

func main() {
	// 初始化数据库
	initDB()

	// 获取所有支部code
	codes := listCode()

	// 设置支部人数
	wg.Add(1)
	go setNum(codes[:len(codes)/2])

	wg.Add(1)
	go setNum(codes[len(codes)/2:])

	wg.Wait()
	fmt.Printf("=== main 结束")
}
