package main

import (
	"fmt"
	"math/rand"
)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] < arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min_ := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min_] > arr[j] {
				min_ = j
			}
		}
		arr[i], arr[min_] = arr[min_], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

// странный вариант быстрой сортировки
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int
	for _, num := range arr[1:] {
		if num < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quickSort(less), pivot)
	result = append(result, quickSort(greater)...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}

func main() {
	l := 100
	arr := createArray(l)
	fmt.Println(arr)

	emptyArr := make([]int, l)
	copy(emptyArr, arr)
	//fmt.Println(emptyArr)
	fmt.Println(bubbleSort(emptyArr))

	copy(emptyArr, arr)
	//fmt.Println(emptyArr)
	fmt.Println(selectionSort(emptyArr))

	copy(emptyArr, arr)
	//fmt.Println(emptyArr)
	fmt.Println(insertionSort(emptyArr))

	copy(emptyArr, arr)
	//fmt.Println(emptyArr)
	fmt.Println(quickSort(emptyArr))

	copy(emptyArr, arr)
	//fmt.Println(emptyArr)
	fmt.Println(mergeSort(emptyArr))
}

func createArray(sizeArr int) (arr []int) {
	arr = make([]int, sizeArr)
	for i := 0; i < sizeArr; i++ {
		arr[i] = rand.Intn(sizeArr)
	}
	return
}
