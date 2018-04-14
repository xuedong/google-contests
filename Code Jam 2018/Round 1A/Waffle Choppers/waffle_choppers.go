package main

import (
	"fmt"
	"strings"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		list := make([]int, 4)
		for i := range list {
			fmt.Scan(&list[i])
		}
		fmt.Scan("\n")
		var r int = list[0]
		var c int = list[1]
		var h int = list[2]
		var v int = list[3]

		listString := make([]string, r)
		for i := range listString {
			fmt.Scanln(&listString[i])
		}

		var numAt int = 0
		for i := range listString {
			numAt += strings.Count(listString[i], "@")
		}
		var numDiners int = (h + 1) * (v + 1)

		if numAt == 0 {
			fmt.Printf("Case #%d: POSSIBLE\n", testCase)
		} else if numAt%numDiners != 0 {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
		} else if h == 1 && v == 1 {
			var numChocos int = numAt / 4
			var flag bool = false
			for i := 1; i < r; i++ {
				for j := 1; j < c; j++ {
					var num1 int = 0
					for k := 0; k < i; k++ {
						for l := 0; l < j; l++ {
							if listString[k][l] == '@' {
								num1 += 1
							}
						}
					}
					var num2 int = 0
					for k := i; k < r; k++ {
						for l := 0; l < j; l++ {
							if listString[k][l] == '@' {
								num2 += 1
							}
						}
					}
					var num3 int = 0
					for k := 0; k < i; k++ {
						for l := j; l < c; l++ {
							if listString[k][l] == '@' {
								num3 += 1
							}
						}
					}
					var num4 int = 0
					for k := i; k < r; k++ {
						for l := j; l < c; l++ {
							if listString[k][l] == '@' {
								num4 += 1
							}
						}
					}
					if num1 == numChocos && num2 == numChocos && num3 == numChocos && num4 == numChocos {
						flag = true
						break
					}
				}
			}

			if flag {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
			} else {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
			}
		}
	}
}
