package main

import (
	"fmt"
	"math"
)

func DeleteLongKeys(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for v := range m {
		if len(v) > 5 {
			newMap[v] = m[v]
		}
	}
	return newMap
}

func SwapKeysAndValues(m map[string]string) map[string]string {
	newMap := make(map[string]string)
	for v := range m {
		newMap[m[v]] = v
	}
	return newMap
}

func SumOfValuesInMap(m map[int]int) int {
	sum_ := 0
	for v := range m {
		sum_ += m[v]
	}
	return sum_
}

func FindMaxKey(m map[int]int) int {
	max_ := math.MinInt32
	for key := range m {
		max_ = max(max_, key)
	}
	return max_
}

func main() {
	arr := map[int]int{10: 37, 3: 19}
	fmt.Println(FindMaxKey(arr))
	fmt.Println(SumOfValuesInMap(arr))
}
