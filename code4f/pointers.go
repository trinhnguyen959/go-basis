package main

import "fmt"

func main() {
	a := 100.1
	var pointer *float64
	pointer = &a
	fmt.Println(pointer)
	fmt.Println()
	fmt.Printf("%T", pointer)
	fmt.Println()

	b := 123
	p2 := new(int)
	p2 = &b
	fmt.Printf("%T", p2)

	//zero value
	fmt.Println()
	var pt *int
	fmt.Println(pt)

	fmt.Println()
	pt2 := new(int)
	fmt.Println(pt2)

	// array pointer
	array := [4]int{1, 2, 3, 4}
	var pt3 *[4]int
	pt3 = &array
	fmt.Println(pt3)
	fmt.Println()
	fmt.Printf("%T", pt3)
	fmt.Println()

	c := 30
	var pt4 = &c
	applyPointer(pt4)
	fmt.Println(c)
}

func applyPointer(pointer *int) {
	*pointer = 777
}
