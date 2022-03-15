package main

import (
    "fmt"
    "math"
)

func main() {
    var testCases int
    fmt.Scanln(&testCases)
    
    for testCase := 1; testCase <= testCases; testCase++ {
        var V, D float64
        fmt.Scan(&V, &D)
        fmt.Scan("\n")
        
        var res float64
        if math.Floor(9.8*D/(V*V)*100/100) == 1.0 {
            res = 45.0
        } else {
            res = math.Asin(9.8*D/(V*V)) * (90.0/math.Pi)
        }
        
        fmt.Printf("Case #%d: %.9f\n", testCase, res)
    }
}