package main

import (
	"fmt"
)

//Contains tests if a string is included in another
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n int
		fmt.Scan(&n)
		var l int
		fmt.Scan(&l)
		fmt.Scan("\n")

		list := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&list[i])
			fmt.Scan("\n")
		}

		numLetters := make([]int, l)
		colStrings := make([]string, l)
		for j := 0; j < l; j++ {
			numLetters[j] = 0
			colStrings[j] = ""
			col := make([]string, n)
			for i := 0; i < n; i++ {
				col[i] = string(list[i][j])
			}
			for k := 0; k < 26; k++ {
				if Contains(col, string(k+65)) {
					numLetters[j]++
					colStrings[j] += string(k + 65)
				}
			}
			// fmt.Printf(colStrings[j] + "\n")
		}
		var numWays = 1
		for j := 0; j < l; j++ {
			numWays *= numLetters[j]
		}
		// fmt.Printf("%d\n", numWays)

		if l == 1 {
			fmt.Printf("Case #%d: -\n", testCase)
		} else if l == 2 {
			var ok = false
			for l1 := 0; l1 < len(colStrings[0]); l1++ {
				for l2 := 0; l2 < len(colStrings[1]); l2++ {
					if Contains(list, string(colStrings[0][l1])+string(colStrings[1][l2])) == false {
						ok = true
						fmt.Printf("Case #%d: %s\n", testCase, string(colStrings[0][l1])+string(colStrings[1][l2]))
						break
					}
				}
				if ok == true {
					break
				}
			}
			if ok == false {
				fmt.Printf("Case #%d: -\n", testCase)
			}
		} else if numWays <= n {
			fmt.Printf("Case #%d: -\n", testCase)
		}
	}
}
