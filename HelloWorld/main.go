package main

import "fmt"

func main() {
	fmt.Printf("Hello, world or नमस्कार or 你好，世界 or Καλημέρα κόσμε or こんにちは世界\n")
	var i complex64 = 9 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v\n", i)
	s := "hello"
	c := []byte(s) // convert string to []byte type
	c[0] = 'c'
	s2 := string(c) // convert back to string type
	fmt.Printf("String is: %s\n", s2)
}

//https://astaxie.gitbooks.io/build-web-application-with-golang/en/02.1.html
