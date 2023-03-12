package main

//func main() {
//	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
//}
//
//// T is a type parameter that is used like normal type inside the function
//// any is a constraint on type i.e T has to implement "any" interface
//func reverse[T any](s []T) []T {
//	l := len(s)
//	r := make([]T, l)
//
//	for i, ele := range s {
//		r[l-i-1] = ele
//	}
//	return r
//}

//// 使用interface中规定的方法和类型来双重约束泛型的参数
//// 使用泛型自带comparacle约束，判断比较
//type Price int
//
//func (i Price) String() string {
//	return strconv.Itoa(int(i))
//}
//
//type Price2 string
//
//func (i Price2) String() string {
//	return string(i)
//}
//
//type ShowPrice interface {
//	String() string
//	~int | ~string
//}
//
//func ShowPriceList[T ShowPrice](s []T) (ret []string) {
//	for _, v := range s {
//		ret = append(ret, v.String())
//	}
//	return
//}
//
func findFunc[T comparable](s []T, v T) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

//
//func main() {
//	fmt.Println(ShowPriceList([]Price{1, 2}))
//	fmt.Println(ShowPriceList([]Price2{"a", "b"}))
//
//	fmt.Println(findFunc([]int{1, 2, 3, 4, 5, 6}, 5))
//	fmt.Println(findFunc([]string{"dudu", "yiyi", "8号"}, "dudu"))
//}

//// 声明一个泛型slice和泛型函数
//type Slice[T any] []T
//
//func echoSlice[T any](s []T) {
//	for _, v := range s {
//		fmt.Println(v)
//	}
//}
//
//func main() {
//	echoSlice(Slice[int]{1, 2, 3, 4})
//	echoSlice(Slice[string]{"a", "b", "c", "d"})
//}

//// 声明一个泛型map
//type Map[K string, V any] map[K]V
//
//func main() {
//	m1 := Map[string, int]{"name": 1}
//	m1["name"] = 2
//
//	m2 := Map[string, string]{"name": "foo"}
//	m2["name"] = "bar"
//
//	fmt.Println(m1, m2)
//}

//// 声明一个泛型channel
//type C[T any] chan T
//
//func main() {
//	a := make(C[int], 10)
//	a <- 1
//	fmt.Println(<-a)
//
//	b := make(C[string], 10)
//	b <- "foo"
//	fmt.Println(<-b)
//}

//
