package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		mapId := make(map[int]map[int]string)

		for j := 0; j < n; j++ {
			var buff1, buff2, buff3 string
			fmt.Fscan(in, &buff1, &buff2, &buff3)
			if buff1 == "CHANGE" {

				name := buff2
				id, _ := strconv.Atoi(buff3)

				mapId[id][j+1] = name // этот id  в эту секунду был занят

			} else {
				id, _ := strconv.Atoi(buff2)
				time, _ := strconv.Atoi(buff3)
				ans := mapId[id][time]

				if ans == "" {
					fmt.Println(404)
				} else {
					fmt.Println(ans)
				}
			}
		}
	}
}
