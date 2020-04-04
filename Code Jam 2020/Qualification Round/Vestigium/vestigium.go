package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var inputN int
		fmt.Scanln(&inputN)

		var trace int
		numRow := 0
		numCol := 0

		inputM := make([][]int, inputN)
		for i := range inputM {
			inputM[i] = make([]int, inputN)
		}

		for idr := 0; idr < inputN; idr++ {
			for idc := 0; idc < inputN; idc++ {
				fmt.Scan(&inputM[idr][idc])
			}
			trace += inputM[idr][idr]
		}

		for idr := 0; idr < inputN; idr++ {
			for i := 1; i <= inputN; i++ {
				flag := true
				for idc := 0; idc < inputN; idc++ {
					if i == inputM[idr][idc] {
						flag = false
						break
					}
				}
				if flag {
					numRow++
					break
				}
			}
		}

		for idc := 0; idc < inputN; idc++ {
			for i := 1; i <= inputN; i++ {
				flag := true
				for idr := 0; idr < inputN; idr++ {
					if i == inputM[idr][idc] {
						flag = false
						break
					}
				}
				if flag {
					numCol++
					break
				}
			}
		}

		fmt.Printf("Case #%d: %d %d %d\n", testCase, trace, numRow, numCol)
	}
}
