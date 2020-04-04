package main

import (
	"fmt"
)

func overlap(a int, b int, c int, d int) bool {
	return !(b <= c || a >= d)
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var inputN int
		fmt.Scanln(&inputN)

		activities := make([][]int, inputN)
		for i := range activities {
			activities[i] = make([]int, 2)
		}

		for idA := 0; idA < inputN; idA++ {
			fmt.Scan(&activities[idA][0])
			fmt.Scan(&activities[idA][1])
		}

		if inputN == 1 {
			fmt.Printf("Case #%d: C\n", testCase)
		} else if inputN == 2 {
			fmt.Printf("Case #%d: CJ\n", testCase)
		} else {
			flag := true
			output := make([]byte, inputN)
			for idA := 0; idA < inputN; idA++ {
				output[idA] = 'N'
			}
			output[0] = 'C'

			for idA := 0; idA < inputN; idA++ {
				if output[idA] == 'N' {
					output[idA] = 'C'
				}

				for idB := idA + 1; idB < inputN; idB++ {
					countC := 0
					countJ := 0
					for i := 0; i < inputN; i++ {
						if i == idB {
							continue
						}
						if overlap(activities[idB][0], activities[idB][1], activities[i][0], activities[i][1]) {
							if output[i] == 'C' {
								countC++
							} else if output[i] == 'J' {
								countJ++
							}
						}
					}

					if countC == 0 && countJ == 0 {
						continue
					} else if countC == 0 && countJ >= 1 {
						output[idB] = 'C'
					} else if countC >= 1 && countJ == 0 {
						output[idB] = 'J'
					} else {
						flag = false
						break
					}
				}

				if !flag {
					break
				}
			}

			if !flag {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
			} else {
				outputStr := string(output)
				fmt.Printf("Case #%d: %s\n", testCase, outputStr)
			}
		}
	}
}
