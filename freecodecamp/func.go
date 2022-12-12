package main

import "fmt"

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("This sum is: ", *s)
	fmt.Println("This sum2 is: ", sum2(2, 3, 4, 5, 6))

	// divide
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// anonymous func
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	//
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()
	fmt.Println("the new name is: ", g.name)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return &result // return pointer
}

func sum2(values ...int) (result int) {
	fmt.Println(values)
	result = 0
	for _, v := range values {
		result += v
	}

	return // not specifics result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
