package main

import (
	"fmt"
	"math"
	"strings"
)

// Функции множественный возврат

func CalculateDigitalRoot(n int) int {
	if n < 10 {
		return n
	} else {
		return CalculateDigitalRoot(SumDigitsRecursive(n))
	}
}

func SumDigitsRecursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return SumDigitsRecursive(n/10) + n%10
}

func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	if k1 == k2 {
		return math.NaN(), math.NaN()
	}

	x := (b2 - b1) / (k1 - k2)
	// Подставляем x в одно из уравнений, чтобы найти y
	y := k1*x + b1

	return x, y
}

func CalculateSeriesSum(n int) float64 {
	var count = 1.0
	for i := 2; i <= n; i++ {
		count += 1.0 / float64(i)
	}
	return count
}

func main() {
	println(CalculateDigitalRoot(12345))
}

func IsPowerOfTwoRecursive(N int) {
	if buff(N) {
		println("YES")
	} else {
		println("NO")
	}
}

func buff(N int) bool {
	if N == 1 {
		return true
	} else if N%2 != 0 {
		return false
	} else if buff(N / 2) {
		return true
	} else {
		return false
	}
}

func add(a, b float64) float64 {
	return a + b
}

func multiply(a, b float64) float64 {
	return a * b
}

func PrintNumbersAscending(n int) {
	var result []string
	for i := 1; i <= n; i++ {
		result = append(result, fmt.Sprintf("%d", i))
	}
	fmt.Println(strings.Join(result, " "))
}
