package main

import (
	"fmt"
	"math"
)

func linearSearch(arr []int, s int) int {
	for i, v := range arr {
		if s == v {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, s int) int {
	var leftPointer = 0
	var rightPointer = len(arr) - 1
	for leftPointer <= rightPointer {
		var midPointer = int(leftPointer+rightPointer) / 2
		var midValue = arr[midPointer]

		if midValue == s {
			return midPointer
		} else if midValue < s {
			leftPointer = midPointer + 1
		} else {
			rightPointer = midPointer - 1
		}
	}
	return -1
}

func jumpSearch(arr []int, s int) int {
	var blockSize = int(math.Sqrt(float64(len(arr))))
	var i = 0
	for {
		if arr[i] >= s {
			break
		}

		if i > len(arr) {
			break
		}

		i += blockSize
	}

	for j := i; j > 0; j-- {
		if arr[j] == s {
			return j
		}
	}
	return -1
}

func main() {
	var n = []int{1, 2, 56, 13, 14, 67, 34, 56}
	f := 14

	fmt.Println(linearSearch(n, f))
	fmt.Println(binarySearch(n, f))
}
