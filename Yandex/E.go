package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

// Строки Руны

func permuteHelper(chars []rune, l int, result *[]string) {
	if l == len(chars)-1 {
		*result = append(*result, string(chars))
	} else {
		for i := l; i <= len(chars)-1; i++ {
			chars[l], chars[i] = chars[i], chars[l]
			permuteHelper(chars, l+1, result)
			chars[l], chars[i] = chars[i], chars[l]
		}
	}
}

func Permutations(input string) []string {
	var result []string
	chars := []rune(input)
	permuteHelper(chars, 0, &result)
	sort.Strings(result)
	return result
}

func IsPalindrome(input string) bool {
	newInput := strings.Join(strings.Split(input, " "), "")
	//fmt.Println(newInput)
	for i := 0; i < len(newInput)/2; i++ {
		if newInput[i] != newInput[len(newInput)-1-i] {
			return false
		}
	}
	return true
}

func isLatin(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}

func ConcatenateStrings(str1, str2 string) string {
	return str1 + " " + str2
}

func StringLength(input string) int {
	return len(input)
}

func main() {
	fmt.Println(StringLength("Hello World"))
	fmt.Println(StringLength(""))
	fmt.Println(ConcatenateStrings("123", "123"))
	fmt.Println(IsPalindrome("a b c cb a"))
}
