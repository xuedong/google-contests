package main

import (
	"fmt"
)

func print5(a int, b int, c int, d int, e int) {
	fmt.Printf("%d %d %d %d %d\n", a, e, b, c, d)
	fmt.Printf("%d %d %d %d %d\n", b, a, d, e, c)
	fmt.Printf("%d %d %d %d %d\n", c, b, a, d, e)
	fmt.Printf("%d %d %d %d %d\n", d, c, e, b, a)
	fmt.Printf("%d %d %d %d %d\n", e, d, c, a, b)
}

func print5d(a int, b int, c int, d int, e int) {
	fmt.Printf("%d %d %d %d %d\n", a, b, c, d, e)
	fmt.Printf("%d %d %d %d %d\n", e, a, b, c, d)
	fmt.Printf("%d %d %d %d %d\n", d, e, a, b, c)
	fmt.Printf("%d %d %d %d %d\n", c, d, e, a, b)
	fmt.Printf("%d %d %d %d %d\n", b, c, d, e, a)
}

func main() {
	var testCases int
	fmt.Scanln(&testCases)

	for testCase := 1; testCase <= testCases; testCase++ {
		var inputN int
		var inputK int
		fmt.Scanln(&inputN, &inputK)

		if inputN == 2 {
			if inputK == 3 {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
			} else if inputK == 2 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 2)
				fmt.Println(2, 1)
			} else if inputK == 4 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(2, 1)
				fmt.Println(1, 2)
			}
		} else if inputN == 3 {
			if inputK%3 != 0 {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
			} else if inputK == 3 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 2, 3)
				fmt.Println(3, 1, 2)
				fmt.Println(2, 3, 1)
			} else if inputK == 6 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(2, 1, 3)
				fmt.Println(3, 2, 1)
				fmt.Println(1, 3, 2)
			} else if inputK == 9 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(3, 2, 1)
				fmt.Println(1, 3, 2)
				fmt.Println(2, 1, 3)
			}
		} else if inputN == 4 {
			if inputK == 5 || inputK == 15 {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
			} else if inputK == 4 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 2, 3, 4)
				fmt.Println(4, 1, 2, 3)
				fmt.Println(3, 4, 1, 2)
				fmt.Println(2, 3, 4, 1)
			} else if inputK == 6 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 2, 3, 4)
				fmt.Println(2, 1, 4, 3)
				fmt.Println(3, 4, 2, 1)
				fmt.Println(4, 3, 1, 2)
			} else if inputK == 7 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 2, 3, 4)
				fmt.Println(3, 1, 4, 2)
				fmt.Println(4, 3, 2, 1)
				fmt.Println(2, 4, 1, 3)
			} else if inputK == 8 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 3, 2, 4)
				fmt.Println(3, 1, 4, 2)
				fmt.Println(2, 4, 3, 1)
				fmt.Println(4, 2, 1, 3)
			} else if inputK == 9 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(1, 3, 4, 2)
				fmt.Println(3, 2, 1, 4)
				fmt.Println(2, 4, 3, 1)
				fmt.Println(4, 1, 2, 3)
			} else if inputK == 10 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(2, 3, 1, 4)
				fmt.Println(3, 2, 4, 1)
				fmt.Println(1, 4, 3, 2)
				fmt.Println(4, 1, 2, 3)
			} else if inputK == 11 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(2, 4, 1, 3)
				fmt.Println(3, 2, 4, 1)
				fmt.Println(4, 1, 3, 2)
				fmt.Println(1, 3, 2, 4)
			} else if inputK == 12 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(2, 4, 3, 1)
				fmt.Println(4, 2, 1, 3)
				fmt.Println(3, 1, 4, 2)
				fmt.Println(1, 3, 2, 4)
			} else if inputK == 13 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(2, 4, 1, 3)
				fmt.Println(4, 3, 2, 1)
				fmt.Println(3, 1, 4, 2)
				fmt.Println(1, 2, 3, 4)
			} else if inputK == 14 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(3, 4, 2, 1)
				fmt.Println(4, 3, 1, 2)
				fmt.Println(2, 1, 4, 3)
				fmt.Println(1, 2, 3, 4)
			} else if inputK == 16 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				fmt.Println(4, 3, 2, 1)
				fmt.Println(1, 4, 3, 2)
				fmt.Println(2, 1, 4, 3)
				fmt.Println(3, 2, 1, 4)
			}
		} else if inputN == 5 {
			if inputK == 6 || inputK == 24 {
				fmt.Printf("Case #%d: IMPOSSIBLE\n", testCase)
			} else if inputK == 5 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5d(1, 2, 3, 4, 5)
			} else if inputK == 7 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(1, 2, 3, 4, 5)
			} else if inputK == 8 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(2, 1, 3, 4, 5)
			} else if inputK == 9 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(1, 3, 2, 4, 5)
			} else if inputK == 10 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5d(2, 3, 4, 5, 1)
			} else if inputK == 11 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(1, 4, 2, 3, 5)
			} else if inputK == 12 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(2, 3, 1, 4, 5)
			} else if inputK == 13 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(3, 2, 1, 4, 5)
			} else if inputK == 14 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(2, 4, 1, 3, 5)
			} else if inputK == 15 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5d(3, 4, 5, 1, 2)
			} else if inputK == 16 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(2, 5, 1, 3, 4)
			} else if inputK == 17 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(3, 4, 1, 2, 5)
			} else if inputK == 18 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(4, 3, 1, 2, 5)
			} else if inputK == 19 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(3, 5, 1, 2, 4)
			} else if inputK == 20 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5d(4, 5, 1, 2, 3)
			} else if inputK == 21 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(5, 3, 1, 2, 4)
			} else if inputK == 22 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(4, 5, 1, 2, 3)
			} else if inputK == 23 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5(5, 4, 1, 2, 3)
			} else if inputK == 25 {
				fmt.Printf("Case #%d: POSSIBLE\n", testCase)
				print5d(5, 1, 2, 3, 4)
			}
		}
	}
}
