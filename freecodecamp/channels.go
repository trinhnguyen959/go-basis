package main

import (
	"fmt"
	"sync"
)

var wgc = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wgc.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wgc.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wgc.Done()
	}(ch)
	wgc.Wait()
}
