package main

import (
	"fmt"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var r int
        var c int
		fmt.Scanln(&r, &c)
        
        n := 2*r+1
        m := 2*c+1

		fmt.Printf("Case #%d:\n", testCase)
        for i := 1; i <= n; i++ {
            for j := 1; j <= m; j++ {
                if i == 1 && j == 1 || i == 1 && j == 2 || i == 2 && j == 1 {
                    fmt.Printf(".")
                } else if i % 2 == 1 && j % 2 == 1 {
                    fmt.Printf("+")
                } else if i % 2 == 1 && j % 2 == 0 {
                    fmt.Printf("-")
                } else if i % 2 == 0 && j % 2 == 0 {
                    fmt.Printf(".")
                } else {
                    fmt.Printf("|")
                }
            }
            fmt.Printf("\n")
        }
	}
}