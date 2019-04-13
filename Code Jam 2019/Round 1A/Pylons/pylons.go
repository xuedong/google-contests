package main

import (
	"fmt"
	"math/rand"
)

// // Pair is a new tuple struct
// type Pair struct {
// 	r int
// 	c int
// }

func count(row int, col int, x int, y int, occupied []int) int {
	sum := 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if i == x || j == y || i-j == x-y || i+j == x+y {
				sum++
			}
		}
	}

	return sum
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var row int
		var col int
		fmt.Scan(&row, &col)

		if row <= 3 && col <= 3 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
		} else if row == 2 && col == 4 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
		} else if row == 4 && col == 2 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
		} else {
			fmt.Printf("Case #%d: POSSIBLE\n", testCase)

			length := row * col
			found := false

			for found == false {
				start := rand.Intn(length)
				path := make([]int, length)
				path[0] = start
				possible := true

				for possible == true {

				}
			}
		}
	}
}
