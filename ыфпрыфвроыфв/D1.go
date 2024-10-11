package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func solve4(source, check string) string {
	if len(check) == 0 || check[0] == ' ' || check[len(check)-1] == ' ' {
		return "no"
	}

	var space bool

	for _, char := range []rune(check) {
		if char == ' ' {
			if space {
				return "no"
			}
			space = true
			continue
		}

		space = false
	}

	var sourceInts []string

	for _, val := range strings.Fields(source) {
		sourceInts = append(sourceInts, val)
	}

	var checkInts []string

	for _, val := range strings.Fields(check) {
		checkInts = append(checkInts, val)
		if len(val) > 1 && val[0] == '0' {
			return "no"
		}
		if len(val) > 1 {
			_, err := strconv.Atoi(val[:2])
			if err != nil {
				return "no"
			}
		}
		if len(val) == 1 {
			_, err := strconv.Atoi(val)
			if err != nil {
				return "no"
			}
			continue
		}
		for _, c := range []rune(val[1:]) {
			if c < '0' || c > '9' {
				return "no"
			}
		}
	}

	if len(sourceInts) != len(checkInts) {
		return "no"
	}

	sort.Strings(sourceInts)

	for i, sourceInt := range sourceInts {
		if sourceInt != checkInts[i] {
			return "no"
		}
	}

	return "yes"
}

func main() {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	s := bufio.NewScanner(os.Stdin)
	defer out.Flush()

	var t int
	s.Scan()
	tStr := s.Text()
	t, _ = strconv.Atoi(tStr)

	var res []string
	for i := 0; i < t; i++ {
		s.Scan()
		_ = s.Text()
		s.Scan()
		s1 := s.Text()
		s1 = strings.TrimRight(s1, "\n")
		s.Scan()
		s2 := s.Text()
		s2 = strings.TrimRight(s2, "\n")
		res = append(res, solve4(s1, s2))
	}

	for _, r := range res {
		fmt.Println(r)
	}
}
