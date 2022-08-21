package main

import (
	"bufio"
	"fmt"
	"os"
)

func isValid(n int, i int, j int) bool {
	return i >= 1 && i <= n && j >= 1 && j <= n
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(n int, i int, j int, direction int, minutes int, coins int, operations []string, values []int) int {
	x, y := i, j
	if direction == 1 {
		y++
	} else if direction == 2 {
		x++
	} else if direction == 3 {
		x--
	} else if direction == 4 {
		y--
	}

	if !isValid(n, x, y) {
		return -10000
	}

	if minutes == 0 {
		operation, value := operations[direction], values[direction]
		if operation == "+" {
			coins += value
		} else if operation == "-" {
			coins -= value
		} else if operation == "*" {
			coins *= value
		} else if operation == "/" {
			if coins < 0 {
				if coins%value == 0 {
					coins /= value
				} else {
					coins = coins/value - 1
				}
			} else {
				coins /= value
			}
		}
		return coins
	}

	ans := dfs(n, x, y, 0, minutes-1, coins, operations, values)
	ans = max(ans, dfs(n, x, y, 1, minutes-1, coins, operations, values))
	ans = max(ans, dfs(n, x, y, 2, minutes-1, coins, operations, values))
	ans = max(ans, dfs(n, x, y, 3, minutes-1, coins, operations, values))
	ans = max(ans, dfs(n, x, y, 4, minutes-1, coins, operations, values))
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var testCases int
	fmt.Fscan(in, &testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var n, p, m, ar, ac int
		fmt.Fscan(in, &n, &p, &m, &ar, &ac)

		operations := make([]string, 5)
		values := make([]int, 5)
		for i := 1; i <= 4; i++ {
			fmt.Fscan(in, &operations[i], &values[i])
		}

		ans := -10000
		for i := 0; i <= 4; i++ {
			ans = max(ans, dfs(n, ar, ac, i, m, 0, operations, values))
		}

		fmt.Printf("Case #%d: %d\n", testCase, ans)
	}
}
