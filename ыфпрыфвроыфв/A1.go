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
		var s string
		fmt.Fscan(in, &s)

		l := len(s)
		if l == 1 {
			fmt.Fprintln(out, "0")
		} else {
			flag := true
			prev := s[0]
			//fmt.Fprintln(out, string(prev))
			for j := 1; j < l; j++ {
				//fmt.Fprintln(out, prev, s[j])
				if s[j] > prev {
					fmt.Fprintln(out, s[:j-1]+s[j:])
					flag = false
					break
				}
				prev = s[j]
			}
			if flag {
				fmt.Fprintln(out, s[:l-1])
			}
		}
	}

}
