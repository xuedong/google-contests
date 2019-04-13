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

// Contains tests if an element belongs to an int slice
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Count counts the number of unavailable cells in a grid
func Count(row int, col int, x int, y int, occupied []int) (int, []int) {
	var availableMoves []int
	sum := 0
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			flattened := (i-1)*col + j
			if i == x || j == y || i-j == x-y || i+j == x+y ||
				Contains(occupied, flattened) {
				sum++
			} else {
				availableMoves = append(availableMoves, flattened)
			}
		}
	}

	return sum, availableMoves
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
				current := rand.Intn(length) + 1
				path := make([]int, length)
				path[0] = current
				pathLength := 0
				var occupied []int
				possible := true

				for possible == true {
					occupied = append(occupied, current)
					y := current % col
					if y == 0 {
						y = col
					}
					x := current/col + 1
					_, availableMoves := Count(row, col, x, y, occupied)
					if len(availableMoves) == 0 {
						possible = false
					} else {
						current = availableMoves[rand.Intn(len(availableMoves))]
						pathLength++
						path[pathLength] = current
					}
				}

				if pathLength == length-1 {
					for i := 0; i < length; i++ {
						y := path[i] % col
						if y == 0 {
							y = col
						}
						x := path[i]/col + 1
						fmt.Printf("%d %d\n", x, y)
					}
					found = true
				}
			}
		}
	}
}
