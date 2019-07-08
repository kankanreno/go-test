package main

import (
	"flag"
	"fmt"
)

//定义一个字符串变量，并制定默认值以及使用方式
var str = flag.String("str", "foo", "the str we are studying")

//定义一个int型字符
var n  = flag.Int("n", 1, "ins n")

func main() {
	// 上面定义了两个简单的参数，在所有参数定义生效前，需要使用flag.Parse()来解析参数
	flag.Parse()
	// 测试上面定义的函数
	fmt.Println("a string flag:",string(*str))
	fmt.Println("ins num:",rune(*n))
}