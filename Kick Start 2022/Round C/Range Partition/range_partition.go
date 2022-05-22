package main

import (
    "fmt"
)

func main() {
    var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
	    var n, x, y int64
	    fmt.Scanln(&n, &x, &y)

	    numerator := x * n * (n+1)
	    denominator := 2 * (x + y)

	    if (numerator % denominator != 0) {
	        fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
	        continue
	    }

	    results := make([]int64, 0)
	    target := numerator / denominator
	    for i := n; i >= 1; i-- {
	        if i <= target {
	            target -= i
	            results = append(results, i)
	        }

	        if target == 0 {
	            break
	        }
	    }

	    fmt.Printf("Case #%d: POSSIBLE\n", testCase)
	    fmt.Printf("%d\n", len(results))
	    for _, num := range results {
	        fmt.Printf("%d ", num)
	    }
	    fmt.Println()
	}
}
