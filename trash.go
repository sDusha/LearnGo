package main

import (
	"fmt"
	"sort"
)

// Рекурсивная функция для генерации перестановок
func permuteHelper(chars []rune, l int, r int, result *[]string) {
	if l == r {
		*result = append(*result, string(chars))
	} else {
		for i := l; i <= r; i++ {
			chars[l], chars[i] = chars[i], chars[l] // меняем местами символы
			permuteHelper(chars, l+1, r, result)    // рекурсивно вызываем с изменённым началом
			chars[l], chars[i] = chars[i], chars[l] // возвращаем символы на место
		}
	}
}

func Permutations(input string) []string {
	var result []string
	chars := []rune(input)
	permuteHelper(chars, 0, len(chars)-1, &result)
	sort.Strings(result)
	return result
}

func main() {
	// Пример использования
	word := "abcd"
	perms := Permutations(word)
	fmt.Println(perms)
}
