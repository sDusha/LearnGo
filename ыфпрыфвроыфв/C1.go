package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func copyMap(m map[int]string) map[int]string {
	new_m := make(map[int]string, len(m))
	for id, v := range m {
		new_m[id] = v
	}
	return new_m
}

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

		currPos := make(map[string]int)
		timeDict := make(map[int]map[int]string)
		for j := 0; j < n; j++ {
			var buff1, buff2, buff3 string
			fmt.Fscan(in, &buff1, &buff2, &buff3)

			if buff1 == "CHANGE" {
				prevTimeDict := make(map[int]string)
				if j > 0 {
					prevTimeDict = copyMap(timeDict[j-1])
				}
				name := buff2
				id, _ := strconv.Atoi(buff3)

				prevID := currPos[name]

				if prevID != id {
					prevTimeDict[prevID] = ""
				}

				currPos[name] = id
				prevTimeDict[id] = name
				timeDict[j] = prevTimeDict
			} else {
				timeDict[j] = copyMap(timeDict[j-1])

				//fmt.Println(timeDict)
				id, _ := strconv.Atoi(buff2)
				time, _ := strconv.Atoi(buff3)
				ans, _ := timeDict[time-1][id]
				if ans == "" {
					fmt.Println(404)
				} else {
					fmt.Println(ans)
				}
			}
		}
	}
}
