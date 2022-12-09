package main

import "fmt"

func main() {
	n := true
	fmt.Printf("%v, %T\n", n, n)

	var m bool // in go, every variable has default value
	fmt.Println(m)

	var us uint16 = 42
	fmt.Printf("%v, %T\n", us, us)

	var a int = 10
	var b int8 = 3

	fmt.Println(a + int(b))

	// so bit
	fmt.Println("\nso bit")
	y := 10
	z := 3

	fmt.Println(y & z)
	fmt.Println(y | z)
	fmt.Println(y ^ z)
	fmt.Println(y &^ z)

	fmt.Println("\nso mu")
	a = 8
	fmt.Println(a << 3)
	fmt.Println(a >> 3)

	fmt.Println("\nrunce")
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
