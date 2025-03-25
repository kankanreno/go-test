package intf

// pkg/interfaces/service.go
type OrderIntf interface {
	Create(uid uint) error
}

type UserIntf interface {
	Speek(uid uint)
}
