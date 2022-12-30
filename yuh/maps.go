package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	number   int
	name     string
	isMale   bool
	subjects []string
}

type People struct {
	name string
	age  int
}

// Student1 /*embedded struct*/
type Student1 struct {
	People
	number  int `Maximum can't over 100` // tag
	subject string
}

func main() {
	studentNameAgeMap := map[string]int{
		"Yuh":  19,
		"Tom":  35,
		"Mike": 40,
	}
	studentNameAgeMap["David"] = 25
	delete(studentNameAgeMap, "Mike")
	fmt.Println(studentNameAgeMap["David"])
	fmt.Println(studentNameAgeMap)
	_, isExist := studentNameAgeMap["David"]
	fmt.Println(isExist)
	fmt.Println(len(studentNameAgeMap))

	// su dung ham make -> slice ko dc lam key trong map
	fmt.Println()
	studentNameAgeMap2 := make(map[string]int)
	fmt.Println(studentNameAgeMap2)

	copyMap := studentNameAgeMap
	copyMap["Yuh"] = 20
	fmt.Println(studentNameAgeMap)
	fmt.Println(copyMap)
	fmt.Println()

	fmt.Println("=============struct==================")
	studentYuh := Student{
		number: 19,
		name:   "Yuh",
		isMale: true,
		subjects: []string{
			"Math",
			"English",
			"Computer",
		},
	}

	fmt.Println(studentYuh.isMale)

	// inner struct
	studentU := struct {
		name string
	}{name: "Yu"}

	copyU := &studentU
	copyU.name = "Copy"
	fmt.Println(studentU)
	fmt.Println(*copyU)

	studentEx := Student1{}
	studentEx.name = "US"
	studentEx.age = 19
	studentEx.number = 3
	studentEx.subject = "this"
	fmt.Println(studentEx)

	t := reflect.TypeOf(studentEx)
	field, _ := t.FieldByName("number")
	fmt.Println(field)
}
