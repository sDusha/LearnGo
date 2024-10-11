package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {
	runes1 := []rune(strings.ToLower(str1))
	runes2 := []rune(strings.ToLower(str2))
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })
	if slices.Equal(runes2, runes1) {
		return true
	}
	return false
}

func SumTwoIntegers(a, b string) (int, error) {
	result1, err1 := strconv.Atoi(a)
	result2, err2 := strconv.Atoi(b)
	if err1 != nil || err2 != nil {
		return -1, errors.New("invalid input, please provide two integers")
	}
	return result1 + result2, nil
}

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}

	return strconv.FormatInt(int64(num), 2), nil
}

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}

	return result, nil
}

func GetCharacterAtPosition(str string, position int) (rune, error) {

	runes := []rune(str)

	if position < 0 || position >= len(runes) {
		return rune(0), errors.New("position out of range")
	}

	return runes[position], nil
}

func DivideIntegers(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return float64(a) / float64(b), nil
}

func ConcatStringsAndInt(str1, str2 string, num int) string {
	result := str1 + str2 + strconv.Itoa(num)
	return result
}

func main() {
	//Ошибки в GO
	fmt.Println(ConcatStringsAndInt("123", "456", 123))
	fmt.Println(DivideIntegers(12, 0))
	fmt.Println(GetCharacterAtPosition("Оле ола вперёд Спартак Москва", 0))
	fmt.Println(GetCharacterAtPosition("Оле ола вперёд Спартак Москва", -40))
	fmt.Println(AreAnagrams("фываке", "екаывф"))
	fmt.Println(AreAnagrams("Банка", "Кабан"))
}
