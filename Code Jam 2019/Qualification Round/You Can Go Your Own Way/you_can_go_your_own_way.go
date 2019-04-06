package main

import (
	"bytes"
	"fmt"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var length int
		var lydiaPath string
		fmt.Scan(&length, &lydiaPath)

		var myPath bytes.Buffer

		for i := 0; i < 2*length-2; i++ {
			if lydiaPath[i] == 'E' {
				myPath.Write([]byte("S"))
			} else {
				myPath.Write([]byte("E"))
			}
		}

		var result string
		result = myPath.String()
		fmt.Printf("Case #%d: %s\n", testCase, result)
	}
}
