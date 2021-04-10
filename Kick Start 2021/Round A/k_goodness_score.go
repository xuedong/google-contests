package main

import (
	"fmt"
	"math"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var N int
		var K int
		fmt.Scanln(&N, &K)
		var str string
		fmt.Scanln(&str)

		var count int = 0

		for i := 0; i < N/2; i++ {
			if str[i] != str[N-i-1] {
				count += 1
			}
		}

		fmt.Printf("Case #%d: %d\n", testCase, int(math.Abs(float64(count-K))))
	}
}
