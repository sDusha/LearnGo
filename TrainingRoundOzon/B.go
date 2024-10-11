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
		var n, p int
		fmt.Fscan(in, &n, &p)

		count_ := 0
		for j := 0; j < n; j++ {
			var a int
			fmt.Fscan(in, &a)
			count_ += (a * p) % 100
		}
		result := fmt.Sprintf("%.2f", float64(count_)/100)
		fmt.Fprintln(out, result)
	}
}
