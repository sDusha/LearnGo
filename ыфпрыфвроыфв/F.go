package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x int
	y int
}

type Points []Point

func (p Points) Len() int { return len(p) }
func (p Points) Less(i, j int) bool {
	if p[i].x == p[j].x {
		return p[i].y > p[j].y
	}
	return p[i].x > p[j].x
}
func (p Points) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, m, s int
	fmt.Fscan(in, &t)
	for q := 0; q < t; q++ {
		fmt.Fscan(in, &n)
		box := make([]Point, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &box[j].x)
			fmt.Fscan(in, &box[j].y)
			if box[j].x < box[j].y {
				box[j].x, box[j].y = box[j].y, box[j].x
			}
		}
		sort.Sort(Points(box))

		fmt.Fscan(in, &m)
		pict := make([]Point, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &pict[j].x)
			fmt.Fscan(in, &pict[j].y)
			if pict[j].x < pict[j].y {
				pict[j].x, pict[j].y = pict[j].y, pict[j].x
			}
		}
		sort.Sort(Points(pict))

		if box[0].x >= pict[0].x {
			var bi, pj int
			s = 0
			for bi < n && pj < m {
				yMax := 0
				iBest := bi
				for bi < n && box[bi].x >= pict[pj].x {
					if box[bi].y >= pict[pj].y {
						if yMax < box[bi].y {
							yMax = box[bi].y
							iBest = bi
						}
					}
					bi++
				}
				if yMax == 0 {
					break
				}
				bi = iBest
				s++
				for pj < m && box[bi].y >= pict[pj].y {
					pj++
				}
			}
			if pj != m {
				fmt.Fprintln(out, -1)
			} else {
				fmt.Fprintln(out, s)
			}
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
