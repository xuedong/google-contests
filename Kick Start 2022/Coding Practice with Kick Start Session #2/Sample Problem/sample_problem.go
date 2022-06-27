package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin)

    var testCases int
    fmt.Fscan(in, &testCases)

    for testCase := 1; testCase <= testCases; testCase++ {
        var n, m int
        fmt.Fscan(in, &n, &m)

        candies := 0
        for i := 1; i <= n; i++ {
            var candy int
            fmt.Fscan(in, &candy)
            candies += candy
        }

        ans := candies % m

        fmt.Printf("Case #%d: %d\n", testCase, ans)
    }
}
