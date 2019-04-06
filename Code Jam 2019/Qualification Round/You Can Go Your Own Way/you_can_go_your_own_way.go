package main

import (
	"fmt"
	"strings"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var length int
		var lydiaPath string
		fmt.Scan(&length, &lydiaPath)

		myPath := strings.Builder{}

		for i := 0; i < 2*length-2; i++ {
			if lydiaPath[i] == 'E' {
				myPath.WriteString("S")
			} else {
				myPath.WriteString("E")
			}
		}

		result := myPath.String()
		fmt.Printf("Case #%d: %s\n", testCase, result)
	}
}
