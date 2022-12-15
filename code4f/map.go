package main

import "fmt"

func main() {
	makeMyMap := make(map[string]int)
	fmt.Println(makeMyMap)

	if makeMyMap == nil {
		fmt.Println("makeMyMap = nil")
	} else {
		fmt.Println("makeMyMap != nil")
	}

	fmt.Println()

	var simpleMyMap map[string]int
	fmt.Println(simpleMyMap)

	if simpleMyMap == nil {
		fmt.Println("simpleMyMap = nil")
	}

	// khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
		"key4": 4,
	}
	fmt.Println(myMap2)

	myMap2["key5"] = 5
	myMap2["key4"] = 15
	fmt.Println(myMap2)
	delete(myMap2, "key5")
	fmt.Println(myMap2)
	fmt.Println(len(myMap2))

	// truy cap phan tu trong map
	value := myMap2["key1"]
	fmt.Println(value)

	value2, found := myMap2["key20"]
	if found {
		fmt.Println(value2)
	} else {
		fmt.Println("not found")
	}
}
