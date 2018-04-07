package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var length int
		fmt.Scanln(&length)
		list := make([]int, length)
		for i := range list {
			fmt.Scan(&list[i])
		}

		odd := make([]int, int(math.Ceil(float64(length)/2.0)))
		for i := range odd {
			odd[i] = list[i*2]
		}
		even := make([]int, length/2)
		for i := range even {
			even[i] = list[i*2+1]
		}

		sort.Ints(odd)
		sort.Ints(even)

		var lengthOdd int = len(odd)
		// fmt.Printf("%d\n", lengthOdd)
		var lengthEven int = len(even)
		// fmt.Printf("%d\n", lengthEven)
		var ok bool = true
		var i int = 0
		var idx int
		if lengthOdd > lengthEven {
			for i < lengthEven && ok {
				if even[i] < odd[i] {
					ok = false
					idx = i * 2
				} else if even[i] > odd[i+1] {
					ok = false
					idx = i*2 + 1
				} else {
					i += 1
				}
			}
		} else {
			for i < lengthEven && ok {
				if even[i] < odd[i] {
					ok = false
					idx = i * 2
				} else if i == lengthEven-1 {
					i += 1
				} else if even[i] > odd[i+1] {
					ok = false
					idx = i*2 + 1
				} else {
					i += 1
				}
			}
		}

		if ok {
			fmt.Printf("Case #%d: OK\n", testCase)
		} else {
			fmt.Printf("Case #%d: %d\n", testCase, idx)
		}
	}
}
