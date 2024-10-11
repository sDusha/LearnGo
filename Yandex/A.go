package main

import "fmt"

// Циклы

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 7; i++ {
		fmt.Println(i, "я уже сделал: ", array[i])
	}
	fmt.Println("8 не успел сделать: ", array[7])
	fmt.Println("9 не успел сделать: ", array[8])
}

func SumOfArray(array [6]int) int {
	count := 0
	for _, v := range array {
		count += v
	}
	return count
}

func FindMinMaxInArray(array [10]int) (int, int) {
	_min := array[0]
	_max := array[0]
	for i := 1; i < len(array); i++ {
		if _min > array[i] {
			_min = array[i]
		}
		if _max < array[i] {
			_max = array[i]
		}
	}
	return _min, _max
}

func ThirdElementInArray(array [6]int) int {
	return array[2]
}

func FiveSteps(array [5]int) [5]int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func main() {
	println(CalculateDigitalRoot(12345))
}
