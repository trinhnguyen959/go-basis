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

	// anonymous struct
	var anonymous = struct {
		email string
		age   int
	}{
		email: "trinhnv@gmail.com",
		age:   31,
	}
	fmt.Println(anonymous)

	// pointer tro toi struct
	st5 := &Student{
		789, "robinson",
	}
	fmt.Println()
	fmt.Println(&st5) // print pointer
	fmt.Println((*st5).id)
	fmt.Println(st5.id)

	// struct in struct, nested struct
	student := Students{
		id:   123,
		name: "This",
		info: StudentInfo{
			class: "0101",
			email: "this@gmail.com",
			age:   27,
		},
	}
	fmt.Println()
	fmt.Println(student)
	fmt.Println()

	// compare 2 struct
	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}
	fmt.Println(s1 == s2)

	// zero value
	var student2 Students
	fmt.Println(student2)
}

type struct1 struct {
	id   int
	name string
}

type Student struct {
	id   int
	name string
}

type Students struct {
	id   int
	name string
	info StudentInfo
}

type StudentInfo struct {
	class string
	email string
	age   int
}
