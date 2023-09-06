package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	cash := make([]int, m*2)
	for i := 0; i < m*2; i += 2 {
		var v int
		fmt.Fscan(in, &v)

		cash[i], cash[i+1] = v, v
	}

	stolen := make([]int, 0)

	for i := m*2 - 1; i >= 0; i-- {
		if n >= cash[i] {
			n -= cash[i]
			stolen = append(stolen, cash[i])
		}
	}

	if n == 0 {
		fmt.Fprintln(out, len(stolen))

		sort.Ints(stolen)

		for _, elem := range stolen {
			fmt.Fprint(out, elem, " ")
		}
	} else {
		fmt.Fprintln(out, "-1")
	}
}
