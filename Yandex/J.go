package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}

func main() {
	result, err1 := GetUTFLength([]byte("Hello World!"))
	result1, err2 := GetUTFLength([]byte(""))
	result2, err3 := GetUTFLength([]byte("asd"))
	result3, err4 := GetUTFLength([]byte("фывыф"))
	fmt.Println(result, err1)
	fmt.Println(result1, err2)
	fmt.Println(result2, err3)
	fmt.Println(result3, err4)
}
