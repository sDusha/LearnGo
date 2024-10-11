package main

import (
	"fmt"
	"sort"
)

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

func main() {
	fmt.Println("Это сортировка из библиотеки Sort")
	fmt.Println("n: длина массива\n" +
		"\nf: функция сравнения, возвращающая true или false. В нашем случае эта функция вернет значение," +
		" если элемент заданного индекса массива больше целевого элемента")

	nums := []int{1, 4, 5, 9, 12, 34}
	target := 5
	res := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	fmt.Println(res)
}
