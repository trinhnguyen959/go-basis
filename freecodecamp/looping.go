package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	fmt.Println("\nfor range")
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Println("\nfor range string")
	y := "hello the world"
	for _, v := range y {
		fmt.Println(string(v))
	}
}
