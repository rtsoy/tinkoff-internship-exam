package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	b := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	start := 0
	for start < n && a[start] == b[start] {
		start++
	}

	end := start
	for i := start; i < n; i++ {
		if a[i] != b[i] {
			end++
		}
	}

	for j := start; j < end; j++ {
		for k := j + 1; k < end; k++ {
			if a[k] < a[j] {
				a[j], a[k] = a[k], a[j]
			}
		}
	}

	for j := start; j < n; j++ {
		if a[j] != b[j] {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}
