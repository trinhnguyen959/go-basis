package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	const myConstant = 42
	// const myConstant float64 = math.Sin(1.57) // cannot add at run time
	fmt.Printf("%v, %T\n", myConstant, myConstant)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
}
