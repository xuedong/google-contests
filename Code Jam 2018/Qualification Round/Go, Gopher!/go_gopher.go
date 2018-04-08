package main

import (
	"fmt"
)

func neighbors(x int, y int, board [][]int) int {
	var n int = board[x+1][y+1] + board[x+1][y] +
		board[x+1][y-1] + board[x][y+1] +
		board[x][y-1] + board[x-1][y+1] +
		board[x-1][y] + board[x-1][y-1]

	return n
}

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)
	for testCase := 1; testCase <= testCases; testCase++ {
		var area int
		fmt.Scanf("%d", &area)
		if area == 20 {
			fmt.Printf("%d %d\n", 500, 500)

			board := make([][]int, 1000)
			for i := range board {
				board[i] = make([]int, 1000)
			}

			for i := range board {
				for j := range board {
					board[i][j] = 0
				}
			}

			for {
				var i, j int
				fmt.Scanf("%d %d", &i, &j)
				if board[i][j] == 0 {
					board[i][j] = 1
				}

				if i == 0 && j == 0 {
					break
				} else {
					// fmt.Printf("%d %d\n", 500, 500)
					if neighbors(500, 500, board) < 8 || board[500][500] == 0 {
						fmt.Printf("%d %d\n", 500, 500)
					} else if neighbors(500, 501, board) < 8 {
						fmt.Printf("%d %d\n", 500, 501)
					} else if neighbors(501, 500, board) < 8 {
						fmt.Printf("%d %d\n", 501, 500)
					} else if neighbors(501, 501, board) < 8 {
						fmt.Printf("%d %d\n", 501, 501)
					} else if neighbors(502, 500, board) < 8 {
						fmt.Printf("%d %d\n", 502, 500)
					} else if neighbors(502, 501, board) < 8 {
						fmt.Printf("%d %d\n", 502, 501)
					}
				}
			}
		} else if area == 200 {
			fmt.Printf("%d %d\n", 501, 501)

			board := make([][]int, 1000)
			for i := range board {
				board[i] = make([]int, 1000)
			}

			for i := range board {
				for j := range board {
					board[i][j] = 0
				}
			}

			for {
				var i, j int
				fmt.Scanf("%d %d", &i, &j)
				if board[i][j] == 0 {
					board[i][j] = 1
				}

				if i == 0 && j == 0 {
					break
				} else {
					// fmt.Printf("%d %d\n", 500, 500)
					if neighbors(501, 501, board) < 8 || board[501][501] == 0 {
						fmt.Printf("%d %d\n", 501, 501)
					} else if neighbors(501, 499, board) < 8 || board[501][499] == 0 {
						fmt.Printf("%d %d\n", 501, 499)
					} else if neighbors(499, 501, board) < 8 || board[499][501] == 0 {
						fmt.Printf("%d %d\n", 499, 501)
					} else if neighbors(499, 499, board) < 8 || board[499][499] == 0 {
						fmt.Printf("%d %d\n", 499, 499)
					} else if neighbors(506, 501, board) < 8 || board[506][501] == 0 {
						fmt.Printf("%d %d\n", 506, 501)
					} else if neighbors(506, 499, board) < 8 || board[506][499] == 0 {
						fmt.Printf("%d %d\n", 506, 499)
					} else if neighbors(504, 501, board) < 8 || board[504][501] == 0 {
						fmt.Printf("%d %d\n", 504, 501)
					} else if neighbors(504, 499, board) < 8 || board[504][499] == 0 {
						fmt.Printf("%d %d\n", 504, 499)
					} else if neighbors(511, 501, board) < 8 || board[511][501] == 0 {
						fmt.Printf("%d %d\n", 511, 501)
					} else if neighbors(511, 499, board) < 8 || board[511][499] == 0 {
						fmt.Printf("%d %d\n", 511, 499)
					} else if neighbors(509, 501, board) < 8 || board[509][501] == 0 {
						fmt.Printf("%d %d\n", 509, 501)
					} else if neighbors(509, 499, board) < 8 || board[509][499] == 0 {
						fmt.Printf("%d %d\n", 509, 499)
					} else if neighbors(516, 501, board) < 8 || board[516][501] == 0 {
						fmt.Printf("%d %d\n", 516, 501)
					} else if neighbors(516, 499, board) < 8 || board[516][499] == 0 {
						fmt.Printf("%d %d\n", 516, 499)
					} else if neighbors(514, 501, board) < 8 || board[514][501] == 0 {
						fmt.Printf("%d %d\n", 514, 501)
					} else if neighbors(514, 499, board) < 8 || board[514][499] == 0 {
						fmt.Printf("%d %d\n", 514, 499)
					} else if neighbors(521, 501, board) < 8 || board[521][501] == 0 {
						fmt.Printf("%d %d\n", 521, 501)
					} else if neighbors(521, 499, board) < 8 || board[521][499] == 0 {
						fmt.Printf("%d %d\n", 521, 499)
					} else if neighbors(519, 501, board) < 8 || board[519][501] == 0 {
						fmt.Printf("%d %d\n", 519, 501)
					} else if neighbors(519, 499, board) < 8 || board[519][499] == 0 {
						fmt.Printf("%d %d\n", 519, 499)
					} else if neighbors(526, 501, board) < 8 || board[526][501] == 0 {
						fmt.Printf("%d %d\n", 526, 501)
					} else if neighbors(526, 499, board) < 8 || board[526][499] == 0 {
						fmt.Printf("%d %d\n", 526, 499)
					} else if neighbors(524, 501, board) < 8 || board[524][501] == 0 {
						fmt.Printf("%d %d\n", 524, 501)
					} else if neighbors(524, 499, board) < 8 || board[524][499] == 0 {
						fmt.Printf("%d %d\n", 524, 499)
					} else if neighbors(531, 501, board) < 8 || board[531][501] == 0 {
						fmt.Printf("%d %d\n", 531, 501)
					} else if neighbors(531, 499, board) < 8 || board[531][499] == 0 {
						fmt.Printf("%d %d\n", 531, 499)
					} else if neighbors(529, 501, board) < 8 || board[529][501] == 0 {
						fmt.Printf("%d %d\n", 529, 501)
					} else if neighbors(529, 499, board) < 8 || board[529][499] == 0 {
						fmt.Printf("%d %d\n", 529, 499)
					} else if neighbors(536, 501, board) < 8 || board[536][501] == 0 {
						fmt.Printf("%d %d\n", 536, 501)
					} else if neighbors(536, 499, board) < 8 || board[536][499] == 0 {
						fmt.Printf("%d %d\n", 536, 499)
					} else if neighbors(534, 501, board) < 8 || board[534][501] == 0 {
						fmt.Printf("%d %d\n", 534, 501)
					} else if neighbors(534, 499, board) < 8 || board[534][499] == 0 {
						fmt.Printf("%d %d\n", 534, 499)
					}
				}
			}
		}
	}
}
