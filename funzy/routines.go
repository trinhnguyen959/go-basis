package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}
	wg.Add(len(links))
	for _, link := range links {
		go checkLink(link)
	}
	wg.Wait()
}

func checkLink(link string) {
	res, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " is down!")
		wg.Done()
		return
	}
	fmt.Println(link + " is OK! " + res.Status)
	wg.Done()
}
