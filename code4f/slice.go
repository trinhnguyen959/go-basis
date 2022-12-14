package main

import "fmt"

func main() {
	// khai bao slice
	var slice []int
	fmt.Println(slice)

	// khoi tao slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// tao slice tu array
	var array = [4]int{1, 2, 3, 4}
	slice2 := array[1:3]
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:]
	fmt.Println(slice4)

	slice5 := array[:3]
	fmt.Println(slice5)

	// tao 1 slice tu 1 slice khac
	slice6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	// slice la reference type
	array1 := [5]int{1, 2, 3, 4, 5}
	slice9 := array1[:]
	slice9[0] = 777
	fmt.Println(slice9)
	fmt.Println(array1)

	// len va capacity
	countries := [...]string{"VN", "USA", "CANADA", "FRANCE", "CHINA"}
	slice10 := countries[0:3]
	fmt.Println(slice10)
	fmt.Println(len(slice10)) // so luong phan tu cua slice
	fmt.Println(cap(slice10)) // so luong phan tu underlying array bat dau tu vi tri start khi
	// slice duoc khoi tao

	// slice voi make
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))
	slice11 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	fmt.Println()
	slice12 := make([]int, 2)
	fmt.Println(slice12)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	// append
	fmt.Println()
	slice13 := make([]int, 2)
	slice13 = append(slice13, 3)
	fmt.Println(slice13)

	// copy
	fmt.Println()
	src := []string{"A", "B", "C", "D"}
	dest := make([]string, 2)
	number := copy(dest, src)
	fmt.Println(dest)
	fmt.Println(number)

	// delete item
	src = append(src[:1], src[2:]...)
	fmt.Println(src)
}
