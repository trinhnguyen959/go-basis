package main

import (
	"fmt"
)

func main() {
	// unbuffered channel
	ch := make(chan int)
	go func() {
		ch <- 100
	}()

	fmt.Println(<-ch)
	fmt.Println("done")

	// buffered channel
	ch2 := make(chan int, 2) // 0 mean unbuffered
	ch2 <- 1
	ch2 <- 2
	//close
	close(ch2)

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

	// select trong truong hop nhieu go routines
	queue := make(chan int)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			queue <- i
		}
		done <- true
	}()

	/*	for {
		select {
		case v := <-queue:
			fmt.Println(v)
		case <-done:
			fmt.Println("done")
			return
		}
	}*/

	fmt.Println("===============")
	queue2 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			if i > 5 {
				close(queue2)
				break
			} else {
				queue2 <- i
			}
		}
	}()

	for value := range queue2 {
		fmt.Println(value)
	}
}
