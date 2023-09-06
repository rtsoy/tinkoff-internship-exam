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

	var (
		n   int
		s   int
		res int
	)

	fmt.Fscan(in, &n, &s)

	for i := 0; i < n; i++ {
		var v int

		fmt.Fscan(in, &v)

		if v > res && v < s {
			res = v
		}
	}

	fmt.Fprintln(out, res)
}
