package main

import "fmt"

func main() {
	/*
		tim gia tri vi tri thu n trong day fibonacci
		value n = n-1 + n-2
	*/
	number := 20
	numberOfWorker := 5
	jobs := make(chan int, number)
	results := make(chan int, number)

	for i := 0; i < numberOfWorker; i++ {
		go worker(jobs, results)
	}

	for i := 1; i <= number; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < number; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fibonacci(n)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
