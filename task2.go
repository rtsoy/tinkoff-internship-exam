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

	var s string
	fmt.Fscan(in, &s)

	required := map[rune]int{
		's': 1,
		'h': 1,
		'e': 1,
		'r': 2,
		'i': 1,
		'f': 1,
	}

	letters := make(map[rune]int)

	for _, char := range s {
		letters[char]++
	}

	minPossible := -1

	for letter, count := range required {
		if availableCount, ok := letters[letter]; ok {
			possibleWords := availableCount / count
			if minPossible == -1 || possibleWords < minPossible {
				minPossible = possibleWords
			}
		} else {
			minPossible = 0
			break
		}
	}

	fmt.Fprintln(out, minPossible)
}
