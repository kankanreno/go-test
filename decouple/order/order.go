package order

import (
	"fmt"
	"github.com/kankanreno/go-test/decouple/intf"
)

type Order struct {
	User intf.UserIntf
}

func NewOrder(u intf.UserIntf) *Order {
	return &Order{User: u}
}

func (o *Order) Create(uid uint) {
	o.User.Speek(uid)
	fmt.Println("创建订单用户:", uid)
}
