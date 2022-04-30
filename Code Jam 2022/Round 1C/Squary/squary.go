package main

import (
	"fmt"
	"strconv"
)

func squary(numbers []int64, k int) string {
	s1 := int64(0)
	s2 := int64(0)
	for _, num := range numbers {
		s1 += num
		s2 += num * num
	}

	if k == 1 {
		if s1 == 0 {
			if s2 == 0 {
				return strconv.Itoa(0)
			}
			return "IMPOSSIBLE"
		}

		rest := (s2 - s1*s1) % (2*s1)
		if rest == 0 {
			return strconv.FormatInt(int64((s2 - s1*s1) / (2*s1)), 10)
		}
		return "IMPOSSIBLE"
	}

	n1 := 1 - s1
	n2 := (s2 - 2*s1 + s1*s1) / 2
	return strconv.FormatInt(n1, 10) + " " + strconv.FormatInt(n2, 10)
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n, k int
		fmt.Scanln(&n, &k)

		numbers := make([]int64, n)
		for i := range numbers {
			fmt.Scan(&numbers[i])
		}

		ans := squary(numbers, k)

		fmt.Printf("Case #%d: %v\n", testCase, ans)
	}
}
