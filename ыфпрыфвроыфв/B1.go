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
		var n int
		fmt.Fscan(in, &n)
		l := make([]uint64, 0, n)
		r := make([]uint64, 0, n)
		for k := 0; k < n; k += 1 {
			var buff uint64
			fmt.Fscan(in, &buff)
			l = append(l, buff)
		}
		for k := 0; k < n; k += 1 {
			var buff uint64
			fmt.Fscan(in, &buff)
			r = append(r, buff)
		}
		//fmt.Fprintln(out, l)
		//fmt.Fprintln(out, r)
		var count uint64 = 1
		for k := 0; k < n; k += 1 {
			var buff uint64
			var diff uint64 = r[k] - l[k]
			buffK := uint64(k + 1)
			buff += diff / buffK
			if l[k]%buffK+diff%buffK >= buffK {
				buff++
			}
			if l[k]%buffK == 0 {
				buff++
			}
			//fmt.Fprintln(out, buff, r[k], l[k], k+1)
			count = (count * buff) % (1000000007)
		}
		fmt.Fprintln(out, count)
	}
}
