package user

import (
	"fmt"
	"github.com/kankanreno/go-test/decouple/intf"
)

type User struct {
	order intf.OrderIntf
}

func NewUser(o intf.OrderIntf) *User {
	return &User{order: o}
}

func (h *User) Speek(uid uint) {
	// 实现细节
	h.order.Create(uid) // 调用接口方法
	fmt.Printf("User Speek: %v\n", uid)
}
