package main

import (
	"errors"
	"fmt"
)

//Массивы

func main() {
	test1 := []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}
	result, err := UnderLimit(test1, 3, 5)
	fmt.Println(result, err)

	result, err = UnderLimit(nil, 3, 5)
	fmt.Println(result, err)

	result, err = UnderLimit([]int{}, 3, 5)
	fmt.Println(result, err)

	result, err = UnderLimit([]int{-13, 0, 6}, -5, 1)
	fmt.Println(result, err)

	result, err = UnderLimit([]int{}, 5, -1)
	fmt.Println(result, err)
}

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	var result []int

	if n < 0 {
		return nil, errors.New("n must be greater or equal than 0")
	}
	if nums == nil {
		return nil, errors.New("nums must be not equal nil")
	}

	for _, num := range nums {
		if num < limit {
			result = append(result, num)
			//println(num)
			n -= 1
			if n == 0 {
				return result, nil
			}
		}
	}
	return result, nil
}
