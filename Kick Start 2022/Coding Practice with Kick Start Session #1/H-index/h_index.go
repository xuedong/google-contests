package main

import (
    "fmt"
)

func main() {
    var testCases int
    fmt.Scanln(&testCases)

    for testCase := 1; testCase <= testCases; testCase++ {
        var numberPapers int
        fmt.Scanln(&numberPapers)

        citations := make([]int, numberPapers)
        for i := range citations {
            fmt.Scan(&citations[i])
        }

        buckets := make([]int, numberPapers+1)
        for i := range buckets {
            buckets[i] = 0
        }

        results := make([]int, numberPapers)
        for i := range results {
            if citations[i] >= numberPapers {
                buckets[numberPapers] += 1
            } else {
                buckets[citations[i]] += 1
            }

            var count int
            for j := numberPapers; j >= 0; j-- {
                count += buckets[j]
                if count >= j {
                    results[i] = j
                    break
                }
            }
        }

        fmt.Printf("Case #%d: ", testCase)
        for i := range results {
            fmt.Printf("%d ", results[i])
        }
        fmt.Printf("\n")
    }
}
