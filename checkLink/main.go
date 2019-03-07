package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.ParseRequestURI("https://github.com/golang/tools/blob/master/present/link.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("u--> ", u)
	d, err := url.ParseRequestURI("https://github.com/golang/tools/blob/masterlink.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("d--> ", d)
	r, err := url.ParseRequestURI("/tools/blob/masterlink.go")
	if err != nil {
		panic(err)
	}
	fmt.Println("r--> ", r)
	t, err := url.ParseRequestURI("gaurav")
	if err != nil {
		panic(err)
	}
	fmt.Println("r--> ", t)
}
