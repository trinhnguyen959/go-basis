package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
var wg = sync.WaitGroup{}

	func main() {
		var msg = "Hello"
		wg.Add(1)
		go func(msg string) {
			fmt.Println(msg)
			wg.Done()
		}(msg)
		msg = "G9"
		wg.Wait()
	}
*/

/*var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}*/

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
