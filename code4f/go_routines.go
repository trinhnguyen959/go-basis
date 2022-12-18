package main

import (
	"fmt"
	"sync"
)

func g1() {
	fmt.Println("g1")
	wg.Done()
}
func g2() {
	fmt.Println("g2")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	/*	go g1()
		goroutine := runtime.NumGoroutine()
		fmt.Println(goroutine)
		time.Sleep(time.Second * 1)*/

	// synchronized goroutines -> dam bao chay theo thu tu
	/*
		logic 1
		go g1()
		go g2()

		logic 2
	*/
	fmt.Println("start")
	wg.Add(2)
	go g1()
	go g2()

	wg.Wait()
	fmt.Println("end")
}
