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

	var maxRoadLength int

	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = make([]int, 3)

		for j := 0; j < 3; j++ {
			fmt.Scan(&roads[i][j])

			if j == 2 {
				if roads[i][j] > maxRoadLength {
					maxRoadLength = roads[i][j]
				}
			}
		}
	}

	states := countStates(roads)
	statesLen := statesLenCalc(states)

	var (
		statesCount = len(states)
		x           = 1
	)

	for i := maxRoadLength; i >= 0; i-- {
		testRoads := make([][]int, 0)
		var exit bool

		for _, road := range roads {
			if road[2] < i {
				testRoads = append(testRoads, road)
			}
		}

		testStates := countStates(testRoads)

		if len(testStates) == statesCount {
			testStatesLen := statesLenCalc(testStates)

			for idx, _ := range testStatesLen {
				if testStatesLen[idx] != statesLen[idx] {
					exit = true
					break
				}
			}
		} else {
			break
		}

		x = i

		if exit {
			break
		}
	}

	fmt.Fprintln(out, x+1)
}

func statesLenCalc(states [][]int) []int {
	res := make([]int, len(states))
	for i := 0; i < len(states); i++ {
		res[i] = len(states[i])
	}

	sort.Ints(res)

	return res
}

func countStates(roads [][]int) [][]int {
	states := make([][]int, 0)

	for i := 0; i < len(roads); i++ {
		var found bool

		for j := 0; j < len(states); j++ {
			for k := 0; k < len(states[j]); k++ {
				if states[j][k] == roads[i][1] {
					states[j] = append(states[j], roads[i][0])
					found = true
					break
				}

				if states[j][k] == roads[i][0] {
					states[j] = append(states[j], roads[i][1])
					found = true
					break
				}
			}

			if found {
				break
			}
		}

		if !found {
			states = append(states, []int{roads[i][0], roads[i][1]})
		}
	}

	for idx, state := range states {
		uniqueMap := make(map[int]bool)
		var result []int

		for _, num := range state {
			if !uniqueMap[num] {
				result = append(result, num)
				uniqueMap[num] = true
			}
		}

		states[idx] = result
	}

	return states
}
