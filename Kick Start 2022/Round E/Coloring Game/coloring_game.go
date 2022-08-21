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
		var n int
		fmt.Fscan(in, &n)

		if n%5 == 0 {
			fmt.Printf("Case #%d: %d\n", testCase, n/5)
		} else {
			fmt.Printf("Case #%d: %d\n", testCase, n/5+1)
		}
	}
}
