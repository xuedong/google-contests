package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var testCases int
	fmt.Fscan(in, &testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n int
		fmt.Fscan(in, &n)

		var s string
		fmt.Fscan(in, &s)

		if n == 1 {
			fmt.Printf("Case #%d: %v\n", testCase, s)
			continue
		}

		flag := false
		for i := 0; i < n-1; i++ {
			if isPalindrome(s[0:i+1]) && isPalindrome(s[i+1:]) {
				fmt.Printf("Case #%d: %v\n", testCase, s[0:i+1])
				flag = true
				break
			}
		}
		if !flag {
			fmt.Printf("Case #%d: %v\n", testCase, s)
		}
	}
}
