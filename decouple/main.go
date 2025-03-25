package main

import (
	"github.com/kankanreno/go-test/decouple/order"
	"github.com/kankanreno/go-test/decouple/user"
)

func main() {
	// 初始化模块B
	orderInst := order.NewOrder(nil) // 先创建空接口实例

	// 初始化模块A（注入实现了OrderService的模块B）
	userInst := user.NewUser(orderInst)

	// 将模块A的UserService实现注入到模块B
	orderInst.User = userInst

	// 测试调用
	userInst.Speek(1001)
	orderInst.Create(1001)
}
