package main

import "fmt"

// interface - multiple interface

type Movement interface {
	move()
}

type Animal interface {
	speak()
}

// NextAnimal embedded interface
type NextAnimal interface {
	Movement
	Animal
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("wow wow!")
}

func (d Dog) move() {
	fmt.Println("dog chay bang 4 chan")
}

// generic type empty interface => tu dong implement

func gout(i interface{}) {
	fmt.Println(i)
}

func main() {
	/*	var animal Animal
		animal = Dog{}
		animal.speak()*/
	dog := Dog{}
	var m Movement = dog
	m.move()
	var animal Animal = dog
	animal.speak()

	fmt.Println("====================")
	dogn := Dog{}
	var na NextAnimal = dogn
	na.move()
	na.speak()

	fmt.Println("====================")
	gout(10)
	gout(10.2)
}
