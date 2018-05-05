package main

import (
	"fmt"
)

func min(array []int) (int, int) {
	var ind int = 0
	var min int = array[0]
	for index, value := range array {
		if min > value {
			min = value
			ind = index
		}
	}
	return min, ind
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)
	for testCase := 1; testCase <= testCases; testCase++ {
		var n int
		fmt.Scan(&n)
		fmt.Scan("\n")

		sold := make([]bool, n)
		for j := 0; j < n; j++ {
			sold[j] = false
		}
		count := make([]int, n)
		for j := 0; j < n; j++ {
			count[j] = 0
		}

		for i := 0; i < n; i++ {
			var d int
			fmt.Scan(&d)

			if d == 0 {
				// fmt.Scan("\n")
				fmt.Printf("%d\n", -1)
			} else {
				list := make([]int, d)
				for i := 0; i < d; i++ {
					fmt.Scan(&list[i])
					count[list[i]] += 1
				}
				// fmt.Scan("\n")
				if d == 1 {
					if sold[list[0]] {
						fmt.Printf("%d\n", -1)
					} else {
						sold[list[0]] = true
						fmt.Printf("%d\n", list[0])
					}
				} else {
					customer := make([]int, d)
					for i := 0; i < d; i++ {
						customer[i] = count[list[i]]
					}
					var m int = customer[0]
					var index int = 0
					for ind, e := range customer {
						if e < m {
							m = e
							index = ind
						}
					}
					if sold[list[index]] {
						fmt.Printf("%d\n", -1)
					} else {
						sold[list[index]] = true
						fmt.Printf("%d\n", list[index])
					}
				}
			}
		}
	}
}
