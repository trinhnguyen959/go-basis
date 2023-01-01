package main

import "fmt"

func main() {
	// tinh tong binh phuong cua 1 day so
	var randomNumbers []int
	for i := 0; i <= 1000; i++ {
		randomNumbers = append(randomNumbers, i)
	}
	// generate pipeline
	inputChan := generatePipeline(randomNumbers)

	// fan-out -> 1 in co nhieu out
	c1 := fanout(inputChan)
	c2 := fanout(inputChan)
	c3 := fanout(inputChan)
	c4 := fanout(inputChan)
	c5 := fanout(inputChan)
	c6 := fanout(inputChan)

	// fan-in
	c := fanin(c1, c2, c3, c4, c5, c6)

	sum := 0
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-c
	}
	fmt.Printf("total number of square: %d", sum)
}

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanout(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func fanin(inputChannel ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputChannel {
			for n := range c {
				in <- n
			}
		}
	}()
	return in
}
