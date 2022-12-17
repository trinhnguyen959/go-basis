package main

import "fmt"

func main() {
	st1 := Student{
		id:   123,
		name: "Ryan",
	}

	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)

	st2 := Student{
		456, "Nguyen van a",
	}

	fmt.Println(st2)
	fmt.Println(st2.id)
	fmt.Println(st2.name)

	var st3 Student = struct {
		id   int
		name string
	}{id: 789, name: "Cy"}

	fmt.Println(st3)

	var st4 Student = struct {
		id   int
		name string
	}{101112, "Cy2"}

	fmt.Println(st4)
}

type Student struct {
	id   int
	name string
}
