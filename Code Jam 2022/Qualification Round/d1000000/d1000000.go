package main

import (
	"fmt"
	"sort"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n int
        fmt.Scanln(&n)
        
        numbers := make([]int, n)
        for i := range numbers {
            fmt.Scan(&numbers[i])
        }
        
        // numbers := []int{10, 10, 7, 6, 7, 4, 4, 5, 7, 4}
        
        sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
        
        for idx, value := range numbers {
            if value < n - idx {
                n = value + idx
                numbers = numbers[:n]
            }
        }
        
        fmt.Printf("Case #%d: %d\n", testCase, len(numbers))
	}
}
