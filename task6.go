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

	var n, m int
	fmt.Fscan(in, &n, &m)

	frequency := make(map[int]int)
	for i := 0; i < n; i++ {
		frequency[i+1] = 1
	}

	gangs := make([][]int, n)
	for i := 0; i < n; i++ {
		gangs[i] = []int{i + 1}
	}

	for i := 0; i < m; i++ {
		var v int
		fmt.Fscan(in, &v)

		switch v {
		case 1:
			var x, y int
			fmt.Fscan(in, &x, &y)

			var (
				currentXGang = -1
				currentYGang = -1
			)

			for j := 0; j < len(gangs); j++ {
				for k := 0; k < len(gangs[j]); k++ {
					if gangs[j][k] == x {
						currentXGang = j
					}

					if gangs[j][k] == y {
						currentYGang = j
					}
				}

				if currentXGang != -1 && currentYGang != -1 {
					break
				}
			}

			gangs[currentXGang] = append(gangs[currentXGang], gangs[currentYGang]...)
			gangs[currentYGang] = []int{}

			for _, elem := range gangs[currentXGang] {
				frequency[elem]++
			}

		case 2:
			var x, y int
			fmt.Fscan(in, &x, &y)

			var (
				currentXGang = -1
				currentYGang = -1
			)

			for j := 0; j < len(gangs); j++ {
				for k := 0; k < len(gangs[j]); k++ {
					if gangs[j][k] == x {
						currentXGang = j
					}

					if gangs[j][k] == y {
						currentYGang = j
					}
				}

				if currentXGang != -1 && currentYGang != -1 {
					break
				}
			}

			if currentYGang == currentXGang {
				fmt.Fprintln(out, "YES")
			} else {
				fmt.Fprintln(out, "NO")
			}

		case 3:
			var x int
			fmt.Fscan(in, &x)

			fmt.Fprintln(out, frequency[x])
		}
	}
}
