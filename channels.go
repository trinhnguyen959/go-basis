package main

import (
	"fmt"
	"sync"
)

var wgc = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wgc.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wgc.Done()
	}()

	go func() {
		ch <- 42
		fmt.Println(<-ch)
		wgc.Done()
	}()
	wgc.Wait()
}
