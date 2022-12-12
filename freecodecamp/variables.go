package main

import (
	"fmt"
	"strconv"
)

// ben ngoai thi phai dung var

var name = "trinhnv"

func main() {
	fmt.Println("hello-world")
	/*	var i int
		i = 42*/
	//var i = 42
	/*	i := 42
		var j float32 = 27*/
	k := 99. // 99 int 99. float64
	fmt.Printf("%v, %T\n", k, k)
	fmt.Println(name)

	var (
		actorName   = "this is"
		companyName = "a"
		season      = 11
	)
	fmt.Println(actorName, companyName, season)

	var s int = 42
	// can not s:=42
	s = 142
	fmt.Println(s)

	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
