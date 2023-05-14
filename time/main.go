package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// REF: https://learnku.com/articles/57583	时间相差秒数_Golang 时间操作大全

func main() {
	fmt.Println("当前时间戳 time.Now().Unix()：", time.Now().Unix())
	fmt.Println("time.Now().Year()：", time.Now().Year())
	fmt.Println("time.Now().Month()：", time.Now().Month())
	fmt.Println("time.Now().Month()：", int(time.Now().Month()))
	fmt.Println("time.Now().Day()：", time.Now().Day())
	fmt.Println("time.Now().Weekday()：", time.Now().Weekday())
	fmt.Println("time.Now().Weekday()：", int(time.Now().Weekday()))
	fmt.Println()

	str := "2019-07-07 08:05:02"
	t, _ := time.Parse("2006-01-02 15:04:05", str)
	fmt.Printf("解析时间 Parse %s：%v\n", str, t)
	fmt.Println()

	fmt.Println("格式化显示当前时间: ", time.Now().Format("2006-01-02 03:04:05"))
	fmt.Printf("格式化显示指定时间, t: %s, v: %s\n", str, t.Format("2006-1-2 3:4:5"))
	fmt.Println("time.Now().Weekday()：", t.Weekday())
	fmt.Println("time.Now().Weekday()：", int(t.Weekday()))
	fullTimeStr := time.Now().Format("20060102150405.000")
	fmt.Println("格式化显示当前时间带毫秒: ", strings.Replace(fullTimeStr, ".", "", -1))
	fmt.Println("格式化显示转换时间戳 time.Unix(1564023522, 0).Format：", time.Unix(1564023522, 0).Format("2006-01-02 15:04:05"))
	fmt.Println("time.Unix(1564023522).Unix()：", time.Unix(1564023522, 0).Unix())
	fmt.Println()

	fromAt, _ := time.Parse("2006-01-02 15:04:05", "2021-12-13 11:11:36")
	toAt := time.Now().Local()
	fmt.Println("fromAt：", fromAt)
	fmt.Println("toAt：", toAt)
	hours := int(math.Floor(toAt.Sub(fromAt).Hours()))
	fmt.Println("diff hours：", hours)
	fmt.Println()

	var tn *time.Time
	//t2 := time.Now()
	fmt.Println("tn: ", tn)
	tn = toTimePtr(time.Now())
	fmt.Println("tn: ", tn)
}

func toTimePtr(t time.Time) *time.Time {
	return &t
}
