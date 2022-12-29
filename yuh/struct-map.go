package main

import "fmt"

func main() {
	// khai bao mang khong can khai bao phan tu
	points := [...]int{7, 10, 5, 11, 20}
	fmt.Printf("%v, %T\n", points, points)
	fmt.Println()

	// khai bao truoc phan tu nhung khong add vao mang
	var points1 [3]int
	fmt.Printf("%v, %T\n", points1, points1)
	fmt.Println()

	var points2 [5]string
	points2[0] = "Alex"
	points2[1] = "Tom"
	points2[2] = "Yuh"
	fmt.Printf("%v, %T\n", points2, points2)
	fmt.Printf("%v, %T\n", points2[2], points2[2])
	fmt.Println(len(points2))

	// copy mang hien tai thanh mang moi => khong thay doi mang cu
	a := [3]int{2, 5, 10}
	b := a
	fmt.Println(a)
	b[1] = 20
	fmt.Println(b)
	// refer mang cu sang mang moi -> con tro
	c := [3]int{1, 2, 3}
	d := &c
	fmt.Println(c)
	d[1] = 10
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println()

	// slice la mang co kich thuoc dong
	e := []int{5, 10, 15}
	fmt.Println(len(e))
	fmt.Println(cap(e))
	fmt.Println()

	// thay doi slice
	f := []int{2, 5, 10, 12, 45, 50}
	// thay doi con tro dau va cuoi
	g := f[:]
	// bat dau tu index 3 tro ve sau
	h := f[3:]
	i := f[:5]
	j := f[3:5]
	fmt.Printf("f %v, %v, %v\n", f, len(f), cap(f))
	fmt.Printf("g %v, %v, %v\n", g, len(g), cap(g))
	fmt.Printf("h %v, %v, %v\n", h, len(h), cap(h))
	fmt.Printf("i %v, %v, %v\n", i, len(i), cap(i))
	fmt.Printf("j %v, %v, %v\n", j, len(j), cap(j))
	fmt.Println()

	// hame make khai bao slice
	a1 := make([]int, 10, 20)

	fmt.Printf("a1 %v, %v, %v\n", a1, len(a1), cap(a1))
	a2 := make([]int, 0)
	a2 = append(a2, 1)
	fmt.Println(a2)
	a2 = append(a2, []int{2, 3, 4, 5, 6, 7}...)
	fmt.Println(a2)
	fmt.Println()

	//stack
	stack := []int{1, 2, 3, 4, 5, 6}
	stack1 := stack[:5]
	fmt.Println(stack1)

	//queue
	queue := stack[1:]
	fmt.Println(queue)

	// lay phan tu o giua
	center := append(stack[:2], stack[3:]...)
	fmt.Println(center)
}
