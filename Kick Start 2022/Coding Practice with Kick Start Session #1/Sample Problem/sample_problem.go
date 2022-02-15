package main

import (
    "fmt"
)

func main () {
    var testCases int
    fmt.Scanln(&testCases)

    for testCase := 1; testCase <= testCases; testCase++ {
        var n int
        var m int
        fmt.Scan(&n, &m)
        fmt.Scan("\n")

        var total int
        candies := make([]int, n)
        for i := range candies {
            fmt.Scan(&candies[i])
            total += candies[i]
        }

        fmt.Printf("Case #%d: %d\n", testCase, total%m)
    }
}
