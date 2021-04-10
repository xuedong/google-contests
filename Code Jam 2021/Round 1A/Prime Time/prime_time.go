package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var M int
		fmt.Scanln(&M)
		primes := make([]int, M)
		numbers := make([]int, M)
		for i := range primes {
			fmt.Scan(&primes[i], &numbers[i])
		}

		

		fmt.Printf("Case #%d: %d\n", testCase, ans)
	}
}
