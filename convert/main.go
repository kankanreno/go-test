package main

import (
	"fmt"
	"github.com/Eun/go-convert"
	"github.com/epiclabs-io/elastic"
	"log"
	"reflect"
)

func main() {
	// ================= github.com/Eun/go-convert
	fmt.Println("=== go-convert")
	/*// convert a int to a string
		var s string
		convert.MustConvert(1, &s)
		fmt.Printf("%s\n", s)

		// convert a map into a struct
		type User struct {
			ID   int
			Name string
		}
		var u User
		convert.MustConvert(map[string]string{
			"Name": "Joe",
			"ID":   "10",
		}, &u)
		fmt.Printf("%#v\n", u)

		// convert Id to int and Groups to []int and keep the rest
		m := map[string]interface{}{
			"Id":     0,
			"Groups": []int{},
		}
		// convert a map into well defined map
		convert.MustConvert(
			map[string]interface{}{
				"Id":      "1",
				"Name":    "Joe",
				"Groups":  []string{"3", "6"},
				"Country": "US",
			},
			&m,
		)
		fmt.Printf("%v\n", m)

		// convert a interface slice into well defined interface slice
		// making the first one an integer, the second a string and the third an float
		sl := []interface{}{0, "", 0.0}
		convert.MustConvert([]string{"1", "2", "3"}, &sl)
		fmt.Printf("%v\n", sl)*/

	//convert strs to uints
	var sl []uint
	convert.MustConvert([]string{"1", "2", "3"}, &sl)
	fmt.Printf("%v\n", sl)

	// convert a struct to map
	u := &struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		23,
		"kankan",
	}
	m := &map[string]interface{}{}
	convert.MustConvert(u, m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("\n\n")

	// ===================== elastic
	fmt.Println("=== elastic")
	var uints []uint
	err := elastic.Set(&uints, []string{"1", "2", "3"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(uints)

	// convert a struct to map
	m2 := &map[string]interface{}{}
	elastic.Set(u, m2)
	fmt.Printf("%#v\n", m2)
	fmt.Printf("\n\n")

	// ===================== custom
	fmt.Println("=== custom")
	// convert a struct to map
	m3 := Struct2Map(*u)
	fmt.Printf("%#v\n", m3)
	fmt.Printf("\n\n")
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get("json")
		fmt.Println(tagName)
		if tagName != "" && tagName != "-" {
			data[tagName] = v.Field(i).Interface()
		}
	}
	return data
}
