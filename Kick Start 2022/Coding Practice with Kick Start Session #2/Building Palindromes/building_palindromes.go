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
        var n, q int
        fmt.Fscan(in, &n, &q)

        var str string
        fmt.Fscan(in, &str)

        prefix := make([][]int, 26)
        for row := range prefix {
            prefix[row] = make([]int, n+1)

            for i := 1; i <= n; i++ {
                if str[i-1] == string('A'+row)[0] {
                    prefix[row][i] = prefix[row][i-1] + 1
                } else {
                    prefix[row][i] = prefix[row][i-1]
                }
            }
        }

        ans := 0
        for i := 1; i <= q; i++ {
            var l, r int
            fmt.Fscan(in, &l, &r)

            odds := 0
            for i := 0; i < 26; i++ {
                if (prefix[i][r] - prefix[i][l-1]) % 2 == 1 {
                    odds += 1
                }
            }

            if odds < 2 {
                ans += 1
            }
        }

        fmt.Printf("Case #%d: %d\n", testCase, ans)
    }
}
