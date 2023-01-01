package main

import "fmt"

func main() {
	fmt.Println("start")
	defer panickier()
	fmt.Println("end")
}

func panickier() {
	fmt.Println("Create panic!")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	panic("Chia 1 so la 0")
}
