package main

import (
	"fmt"
	"sync"
)

func main() {
	// demo 2 way channel
	var wg = sync.WaitGroup{}
	ch := make(chan int, 50) // channel with buffer
	wg.Add(2)
	/*	go func() {
			i := <-ch
			fmt.Println(i)
			ch <- 27
			wg.Done()
		}()

		go func() {
			ch <- 42
			fmt.Println(<-ch)
			wg.Done()
		}()
		wg.Wait()*/

	// chanel nhan
	go func(ch <-chan int) {
		/*		for i := range ch {
				fmt.Println(i)
			}*/
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)

	// channel gui
	go func(ch chan<- int) {
		a := 42
		ch <- a
		ch <- a + 1
		ch <- a + 2
		ch <- a + 3
		ch <- a + 4
		ch <- a + 5
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
