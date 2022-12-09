package main

import "fmt"

func main() {
	grades := [...]int{100, 92, 95}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Println()
	for i := 0; i < len(grades); i++ {
		fmt.Printf("grade: %v\n", grades[i])
	}

	fmt.Printf("Length: %v\n", len(grades))
	fmt.Printf("Capacity: %v\n", cap(grades))

	x := []int{1, 2, 3}
	y := x
	y[1] = 3

	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("\n=======Slice=======")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	a[5] = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("\n=====Some builtin function====")
	f := make([]int, 1, 5)
	fmt.Printf("length: %v", len(f))
	fmt.Printf("\ncapacity: %v", cap(f))

	fmt.Println()
	fmt.Println(f)
	fmt.Println("\n=====push====")
	f = append(f, 1, 2, 3, 4, 5)
	fmt.Println(f)
	fmt.Printf("\nlength after add: %v", len(f))
	fmt.Printf("\ncapacity after add: %v", cap(f))

	fmt.Println("\n\n=====pop====")
	g := f[1:]
	//g := f[:len(f)-1]
	fmt.Println(f)
	fmt.Println(g)
	h := append(f[3:], f[:3]...)
	fmt.Println(h)
}
