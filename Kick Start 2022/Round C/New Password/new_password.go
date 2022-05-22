package main

import (
    "fmt"
    "unicode"
)

func nchars(b byte, n int) string {
    s := make([]byte, n)
    for i := 0; i < n; i++ {
        s[i] = b
    }
    return string(s)
}

func main() {
    var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n int
		fmt.Scanln(&n)

		var s string
		fmt.Scanln(&s)

		conditions := make([]int, 4)
		for i := 0; i < n; i++ {
			ch := rune(s[i])
			if unicode.IsUpper(ch) {
			    conditions[0]++
			} else if unicode.IsLower(ch) {
			    conditions[1]++
			} else if unicode.IsDigit(ch) {
			    conditions[2]++
			} else if ch == '#' || ch == '@' || ch == '*' || ch == '&' {
			    conditions[3]++
			}
		}

		if conditions[0] == 0 {
		    s = s + string('A')
		}
		if conditions[1] == 0 {
		    s = s + string('a')
		}
		if conditions[2] == 0 {
		    s = s + string('1')
		}
		if conditions[3] == 0 {
		    s = s + string('#')
		}

		if len(s) < 7 {
		    s = s + nchars('a', 7-len(s))
		}

		fmt.Printf("Case #%d: %v\n", testCase, s)
	}
}
