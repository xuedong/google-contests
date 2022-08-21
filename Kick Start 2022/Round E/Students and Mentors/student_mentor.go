package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func binary_search(nums []int, target int) int {
	left, right := -1, len(nums)

	for right-left > 1 {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid
		} else {
			right = mid
		}
	}

	return left
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var testCases int
	fmt.Fscan(in, &testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n int
		fmt.Fscan(in, &n)

		ratings := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &ratings[i])
		}

		sorted := make([]int, n)
		copy(sorted, ratings)

		fmt.Printf("Case #%d: ", testCase)

		sort.Ints(sorted)
		for i := 0; i < n; i++ {
			target := 2 * ratings[i]
			idx := binary_search(sorted, target)

			if sorted[idx] == ratings[i] {
				if idx == 0 {
					fmt.Printf("%d ", -1)
				} else {
					fmt.Printf("%d ", sorted[idx-1])
				}
			} else {
				fmt.Printf("%d ", sorted[idx])
			}
		}
		fmt.Println()
	}
}
