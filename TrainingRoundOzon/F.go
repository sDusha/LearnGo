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
		var n, k, m int
		fmt.Fscan(in, &n, &k)
		//fmt.Println(n, k)

		fmt.Fscan(in, &m)
		for j := 0; j < m; j++ {
			var buff int
			fmt.Fscan(in, buff)
		}
	}
}
