package main

import (
	"github.com/sirupsen/logrus"
	"reflect"
)

//type Stringer interface {
//	String() string
//}
//
//type Binary uint64
//
//func (i Binary) String() string {
//	return strconv.Uitob64(i.Get(), 2)
//}
//
//func (i Binary) Get() uint64 {
//	return uint64(i)
//}
//
//func main() {
//	b := Binary()
//	s := Stringer(b)
//	fmt.Println(s.String())
//}


type User struct {
	ID uint
	Name string
	Pass string
	Status int
}

func main() {
	logrus.Infof("=== %s", reflect.TypeOf(&User{}).Elem().Name())
}