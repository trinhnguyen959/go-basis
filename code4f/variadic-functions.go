package main

import "fmt"

func main() {
	addItems(1, 100, 200, 300, 400)
	list := []int{1, 2, 3, 4}
	addItems(5, list...)

	change(list...)
	fmt.Println(list)
}

func addItems(item int, list ...int) {
	list = append(list, item, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(list)
}

func change(list ...int) {
	list[0] = 999
}
