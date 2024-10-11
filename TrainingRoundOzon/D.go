package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscanf(in, "%d", &t)
	for i := 0; i < t; i++ {
		var n, buff, number1, number2, maxLength int
		fmt.Fscan(in, &n)
		if n != 1 {
			fmt.Fscan(in, &number1)
			fmt.Fscan(in, &number2)
			pos1 := 0
			pos2 := 1
			length := 2

			for j := 2; j < n; j++ {
				fmt.Fscan(in, &buff)
				if buff == number1 {
					pos1, pos2 = pos2, j
					number1, number2 = number2, number1
					length += 1
				} else if buff == number2 {
					pos2 = j
					length += 1
				} else if number1 == number2 {
					number2 = buff
					pos1, pos2 = pos2, j
					length += 1
				} else {
					maxLength = max(maxLength, length)
					length = j - pos1
					pos1, pos2 = pos2, j
					number1, number2 = number2, buff
				}
				//fmt.Fprintln(out, length)
				//
			}
			fmt.Fprintln(out, max(maxLength, length))
		} else {
			fmt.Fscan(in, &buff)
			fmt.Fprintln(out, 1)
		}

	}
}
