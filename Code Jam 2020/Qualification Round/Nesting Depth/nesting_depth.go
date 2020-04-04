package main

import (
	"bytes"
	"fmt"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var input string
		fmt.Scan(&input)
		inputStr := []rune(input)

		var output bytes.Buffer

		length := len(inputStr)
		if length == 1 {
			numPr := int(inputStr[0] - '0')
			for j := 1; j <= numPr; j++ {
				output.WriteRune('(')
			}
			output.WriteRune(inputStr[0])
			for j := 1; j <= numPr; j++ {
				output.WriteRune(')')
			}
		} else {
			for i := range inputStr {
				if i == 0 {
					numOp := int(inputStr[i] - '0')
					for j := 1; j <= numOp; j++ {
						output.WriteRune('(')
					}
					output.WriteRune(inputStr[i])
				} else {
					numP := int(inputStr[i-1] - '0')
					numN := int(inputStr[i] - '0')
					numPr := numN - numP
					if numPr == 0 {
						output.WriteRune(inputStr[i])
					} else if numPr > 0 {
						for j := 1; j <= numPr; j++ {
							output.WriteRune('(')
						}
						output.WriteRune(inputStr[i])
					} else {
						for j := 1; j <= -numPr; j++ {
							output.WriteRune(')')
						}
						output.WriteRune(inputStr[i])
					}

					if i == length-1 {
						numCl := int(inputStr[i] - '0')
						for j := 1; j <= numCl; j++ {
							output.WriteRune(')')
						}
					}
				}
			}
		}

		outputStr := output.String()

		fmt.Printf("Case #%d: %s\n", testCase, outputStr)
	}
}
