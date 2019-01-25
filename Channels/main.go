package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://google.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}
}

func checkLink(link string, c chan string) {
	time.Sleep(3 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Might be down!")
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}
