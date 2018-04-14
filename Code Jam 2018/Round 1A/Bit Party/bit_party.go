package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var r int
		fmt.Scan(&r)
		var b int
		fmt.Scan(&b)
		var c int
		fmt.Scan(&c)
		fmt.Scan("\n")

		listM := make([]int, c)
		listS := make([]int, c)
		listP := make([]int, c)
		for i := 0; i < c; i++ {
			fmt.Scan(&listM[i])
			fmt.Scan(&listS[i])
			fmt.Scan(&listP[i])
			fmt.Scan("\n")
		}

		var totalLimits int = 0
		for i := 0; i < c; i++ {
			totalLimits += listM[i]
		}

		if totalLimits == b {
			var best int = 0
			for i := 0; i < c; i++ {
				var time int = listS[i]*listM[i] + listP[i]
				if time > best {
					best = time
				}
			}
			fmt.Printf("Case #%d: %d\n", testCase, best)
		} else if totalLimits > b {
			var best int = 1000000
			for i := 0; i < c; i++ {
				var time int = listS[i]*listM[i] + listP[i]
				if time < best {
					best = time
				}
			}
			fmt.Printf("Case #%d: %d\n", testCase, best)
		}
	}
}
