package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	solve(os.Stdin, os.Stdout)
}

func solve(r io.Reader, w io.Writer) {
	var t, n, a, b int
	_, _ = fmt.Fscanf(r, "%d\n", &t)
	results := make([][]int, t)
	for i := 0; i < t; i++ {
		_, _ = fmt.Fscanf(r, "%d\n%d\n%d\n", &n, &a, &b)
		results[i] = make([]int, 0, n)
		for j := 0; j < n; j++ {
			val := j * a
			val += (n - j - 1) * b
			results[i] = addresult(results[i], val)
		}
		sort.Ints(results[i])
	}
	for i := 0; i < t; i++ {
		for j := 0; j < len(results[i]); j++ {
			fmt.Fprintf(w, "%d%s", results[i][j], func() string {
				if j < len(results[i])-1 {
					return " "
				}
				return "\n"
			}())
		}
	}
}

func addresult(r []int, val int) []int {
	for _, v := range r {
		if val == v {
			return r
		}
	}
	return append(r, val)
}
