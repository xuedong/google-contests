package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var total int
		var picked int
		fmt.Scanln(&total, &picked)

		ratings := make([]int, total)
		for i := range ratings {
			fmt.Scan(&ratings[i])
		}

		// sort the rating list
		sort.Ints(ratings)

		// compute the sum of differences between the p-th element and its precedents
		best := math.MaxInt32
		for i := picked - 1; i < total; i++ {
			time := 0
			for j := i - picked + 1; j < i; j++ {
				time += ratings[i] - ratings[j]
			}
			if time < best {
				best = time
			}
		}

		fmt.Printf("Case #%d: %d\n", testCase, best)
	}
}
