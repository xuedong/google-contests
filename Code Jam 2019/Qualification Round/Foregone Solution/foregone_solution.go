package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var inputN int
		fmt.Scan(&inputN)

		var outputA int
		var outputB int

		outputA = inputN
		outputB = 0

		var str = strconv.Itoa(inputN)
		var length = len(str)

		for i := 0; i < length; i++ {
			if str[length-1-i] == '4' {
				outputB += int(math.Pow10(i))
				outputA -= int(math.Pow10(i))
			}
		}

		fmt.Printf("Case #%d: %d %d\n", testCase, outputA, outputB)
	}
}
