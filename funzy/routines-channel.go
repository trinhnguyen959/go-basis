package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLinks(c, link)

	}

	for {
		go checkLinks(c, <-c)

		//fmt.Println(<-c)
	}
}

func checkLinks(c chan string, link string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " is down!")
		c <- link
		return
	}
	fmt.Println(link + " is OK! " + res.Status)
	c <- link
}
