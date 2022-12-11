package main

import "fmt"

func main() {
	/*	a := 42
		b := a
		fmt.Println(a, b)
		a = 27
		fmt.Println(a, b)*/

	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 37
	fmt.Println(a, *b)
}
