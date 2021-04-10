package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var length int
		fmt.Scanln(&length)
		list := make([]int, length)
		for i := range list {
			fmt.Scan(&list[i])
		}

		ans := 0
		for i := 1; i < length; i++ {
			curr := list[i]
			prev := list[i-1]
			currStr := strconv.Itoa(curr)
			prevStr := strconv.Itoa(prev)
			if curr > prev {
				continue
			} else if curr == prev {
				ans += 1
				list[i] *= 10
			} else {
				j := 0
				lenCurr := len(currStr)
				lenPrev := len(prevStr)
				for j < lenCurr {
					if currStr[j] != prevStr[j] {
						break
					}
					j += 1
				}

				if j == lenCurr {
					rest, _ := strconv.Atoi(prevStr[j:lenPrev])
					if rest + 1 == int(math.Pow10(lenPrev-lenCurr)) {
						shift := lenPrev - lenCurr + 1
						ans += shift
						list[i] *= int(math.Pow10(shift))
					} else {
						ans += lenPrev - lenCurr
						list[i] = list[i-1] + 1
					}
				} else if currStr[j] > prevStr[j] {
					shift := lenPrev - lenCurr
					ans += shift
					list[i] *= int(math.Pow10(shift))
				} else {
					shift := lenPrev - lenCurr + 1
					ans += shift
					list[i] *= int(math.Pow10(shift))
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", testCase, ans)
	}
}
