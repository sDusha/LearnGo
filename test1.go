package main

import (
	"fmt"
)

func hello(a string) {
	fmt.Println("Hello World", a)
}

func main() {
	hello("1")
	var (
		name string = "Andrew"
		age  int64  = 27
	)
	fmt.Println(name, age)

	var hello = "Hello World2"
	fmt.Println(hello)
	fmt.Println("Hello, playground")

	flag := true // автоматически определяется тип данных
	fmt.Println(flag)

	var state = "Texas"
	fmt.Println(state)

	const pi float64 = 3.1415
	fmt.Println(pi)

	var m float32 = 10 / 4.0 // 2.5
	fmt.Println(m)

	var numbers = [...]int{1, 2, 3, 4, 5} // длина массива 5
	numbers2 := [...]int{1, 2, 3}         // длина массива 3
	fmt.Println(numbers)                  // [1 2 3 4 5]
	fmt.Println(numbers2)                 // [1 2 3]

	colors := [3]string{2: "blue", 0: "red", 1: "green"}
	fmt.Println(colors[2]) // blue
}
