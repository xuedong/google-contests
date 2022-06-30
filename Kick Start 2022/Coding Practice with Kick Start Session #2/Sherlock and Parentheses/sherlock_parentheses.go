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
        var l, r int
        fmt.Fscan(in, &l, &r)

        var n int
        if l > r {
            n = r
        } else {
            n = l
        }

        ans := n * (n+1) / 2

        fmt.Printf("Case #%d: %d\n", testCase, ans)
    }
}
