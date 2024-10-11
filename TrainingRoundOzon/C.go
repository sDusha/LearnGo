package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		var m, buff, buff1 int
		fmt.Fscan(in, &m)
		arr := make([]int, 0, m)
		arr2 := make([]int, 0, m)
		for k := 0; k < m; k += 2 {
			fmt.Fscan(in, &buff)
			arr = append(arr, buff)
			fmt.Fscan(in, &buff1)
			for j := 0; j < buff1; j++ {
				k++
				var buff2 int
				fmt.Fscan(in, &buff2)
				arr2 = append(arr2, buff2)
			}
		}
		//fmt.Println(arr)
		//fmt.Println(arr2)

		for _, v := range arr {
			if !slices.Contains(arr2, v) {
				fmt.Fprintln(out, v)
			}
		}
	}
}
