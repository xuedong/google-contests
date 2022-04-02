package main

import (
	"fmt"
)

func min(a int, b int) int {
    if a >= b {
        return b
    } else {
        return a
    }
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var c1, m1, y1, k1 int
        fmt.Scanln(&c1, &m1, &y1, &k1)
        var c2, m2, y2, k2 int
        fmt.Scanln(&c2, &m2, &y2, &k2)
        var c3, m3, y3, k3 int
        fmt.Scanln(&c3, &m3, &y3, &k3)
        
        // c1, m1, y1, k1 = 300000, 200000, 300000, 500000
        // c2, m2, y2, k2 = 300000, 200000, 500000, 300000
        // c3, m3, y3, k3 = 300000, 500000, 300000, 200000
        
        mc := min(min(c1, c2), c3)
        mm := min(min(m1, m2), m3)
        my := min(min(y1, y2), y3)
        mk := min(min(k1, k2), k3)
        
        total := mc + mm + my + mk
        if total < 1000000 {
            fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
        } else {
            rest := total - 1000000
            
            res := make([]int, 4)
            res[0], res[1], res[2], res[3] = mc, mm, my, mk
            for i := 0; i <= 3; i++ {
                if rest > 0 {
                    var temp = min(res[i], rest)
                    res[i] -= temp
                    rest -= temp
                }
            }
            
            fmt.Printf("Case #%d: %d %d %d %d\n", testCase, res[0], res[1], res[2], res[3])
        }
	}
}
