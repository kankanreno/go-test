package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	for {
		go func() {
			strs := []string{
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"浏览页面",
				"加入购物车",
				"加入购物车",
				"加入购物车",
				"加入购物车",
				"加入购物车",
				"加入购物车",
				"提交订单",
				"提交订单",
				"提交订单",
				"提交订单",
				"加入收藏",
				"加入收藏",
				"加入收藏",
				"查看订单",
				"查看订单",
				"评论商品",
				"error",
				"error",
				"error",
				"error",
			}
			randomIndex := rand.Intn(len(strs))
			pick := strs[randomIndex]

			now := time.Now()
			date := now.Format("2006-01-02 15:04:05")

			result := date + "|" + pick
			log.Println(result)
			writeLog(result)
		}()

		time.Sleep(1 * time.Second)
	}

}

func writeLog(str string) {
	f, err := os.OpenFile("/var/log/mock/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("open file err: ", err)
	}
	defer f.Close()

	//f.WriteString(`172.0.0.12 - - [04/Mar/2018:13:49:52 +0000] http "GET /foo?query=t HTTP/1.0" 200 2133 "-" "KeepAliveClient" "-" 1.005 1.854 >> access.log` + "\n")
	f.WriteString(str + "\n")
}
