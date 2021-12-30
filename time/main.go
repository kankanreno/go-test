package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("当前时间戳 time.Now().Unix()：", time.Now().Unix())
	fmt.Println("time.Now().Year()：", time.Now().Year())
	fmt.Println("time.Now().Month()：", time.Now().Month())
	fmt.Println("time.Now().Day()：", time.Now().Day())

	t, _ := time.Parse("2006-01-02 15:04:05", "2019-07-25 10:58:42")
	fmt.Println("解析时间 Parse 2019-07-25 10:58:42：", t)

	fmt.Println("格式化显示当前时间: ", time.Now().Format("2006-01-02 15:04:05"))
	fullTimeStr := time.Now().Format("20060102150405.000")
	fmt.Println("格式化显示当前时间带毫秒: ", strings.Replace(fullTimeStr, ".", "", -1))
	fmt.Println("格式化显示转换时间戳 time.Unix(1564023522, 0).Format：", time.Unix(1564023522, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("time.Unix(1564023522).Unix()：", time.Unix(1564023522, 0).Unix())

	fromAt, _ := time.Parse("2006-01-02 15:04:05", "2021-12-13 11:11:36")
	toAt := time.Now().Local()
	fmt.Println("fromAt：", fromAt)
	fmt.Println("toAt：", toAt)
	hours := int(math.Floor(toAt.Sub(fromAt).Hours()))
	fmt.Println("diff hours：", hours)
}
