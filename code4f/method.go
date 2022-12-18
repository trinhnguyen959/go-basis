package main

import "fmt"

type StudentM struct {
	name string
}

func getName(student StudentM) string {
	return student.name
}

/*
	define method
	func (t Type) methodName(params) returns {// body code}
	t Type => receiver
	value and pointer receiver
*/

/*
 1. value receiver

khong lam thay doi gia tri cua struct khi set o method
*/
func (s StudentM) changeName() {
	fmt.Printf("p2= %p\n", &s)
	s.name = "Robin"
}

/* 2. pointer receiver
 */
func (s *StudentM) changeName2() {
	fmt.Printf("p3= %p\n", s)
	s.name = "Robinson"
}

// non struct
type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s%s", s, str)
}

func main() {
	student := StudentM{
		"trinity",
	}
	name := getName(student)
	fmt.Println(name)
	fmt.Printf("p1= %p\n", &student)
	fmt.Println()

	student.changeName()
	fmt.Println(student.name)
	fmt.Println()

	student.changeName2()
	fmt.Println(student.name)

	// none struct
	s1 := String("a")
	newStr := s1.append("b")
	fmt.Println(newStr)
}
