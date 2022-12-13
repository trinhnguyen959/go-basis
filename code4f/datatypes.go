package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	fmt.Println("===================")
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println(bits.OnesCount64(math.MaxUint64))

	fmt.Println("===================")
	var myByte byte = 1
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)
	fmt.Println()
	fmt.Println(math.MaxUint8)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	fmt.Println("========runes=========")
	var myString = "giới định tuệ"
	runes := []rune(myString)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	fmt.Println()
	fmt.Printf("%T", runes)
	fmt.Println()
	fmt.Printf("%T - %U - %d", '🤣', '🤣', '🤣')
	fmt.Println()

	fmt.Println("========type conversion=========")
	a := 1
	b := 2.1
	fmt.Println(float64(a) + b)

	fmt.Println("========const=========")
	const PI = 3.14159
	fmt.Println(PI)
	fmt.Printf("%T", PI)
}
