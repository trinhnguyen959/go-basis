package main

import "fmt"

func main() {
	var a = 12
	var b = &a
	fmt.Println(a, *b)
	a = 24
	fmt.Println(a, *b)
	*b = 32
	fmt.Println(a, *b)

	var arr = []int{1, 2, 3}
	var brr = arr
	arr[1] = 10
	brr[2] = 100
	fmt.Println(arr, brr)
	fmt.Println()

	fmt.Println("==============struct=============")
	var astr *myStruct
	//var astr = new(myStruct) // the same
	fmt.Println(astr)
	astr = &myStruct{
		number: 10,
	}
	astr.number = 20
	fmt.Println(astr)

	fmt.Println("==============maps=============")
	var as = map[string]string{"Ka": "ME", "Me": "KA", "Ha": "HA"}
	var bs = as
	fmt.Println(as, bs)
	as["Me"] = "Ha"
	fmt.Println(as, bs)
}

type myStruct struct {
	number int
}
