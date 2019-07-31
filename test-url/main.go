package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	//u, err := url.Parse("http://localhost:8000/rest/v1/issues_page?order=desc&foo=bar")
	u, err := url.Parse("http://localhost:8000/rest/v1/issues_status_count?")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("u:", u)
	fmt.Println("u.Scheme:", u.Scheme)
	fmt.Println("u.Opaque:", u.Opaque)
	fmt.Println("u.Host:", u.Host)
	fmt.Println("u.Path:", u.Path)
	fmt.Println("u.RawPath:", u.RawPath)
	fmt.Println("u.ForceQuery:", u.ForceQuery)
	fmt.Println("u.RawQuery:", u.RawQuery)
	fmt.Println("u.Fragment:", u.Fragment)

	fmt.Println("u.RequestURI():", u.RequestURI())
}
