package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

		var jsonData string
		for j := 0; j <= n; j++ {
			s, _ := in.ReadString('\n')
			jsonData += strings.TrimSpace(s) // удаляет лишнее
		}
		fmt.Fprintln(out, jsonData)

		jsonObj, _ := json.Marshal(jsonData)
		fmt.Fprintln(out, jsonObj)

	}
}
