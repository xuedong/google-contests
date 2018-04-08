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

			startPoints := [10]int{500, 505, 510, 515, 520,
				525, 530, 535, 540, 545}

			for _, start := range startPoints {
				// fmt.Println(start)
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
						if neighbors(start, 500, board) < 8 || board[start][500] == 0 {
							fmt.Printf("%d %d\n", start, 500)
						} else if neighbors(start, 501, board) < 8 {
							fmt.Printf("%d %d\n", start, 501)
						} else if neighbors(start+1, 500, board) < 8 {
							fmt.Printf("%d %d\n", start+1, 500)
						} else if neighbors(start+1, 501, board) < 8 {
							fmt.Printf("%d %d\n", start+1, 501)
						} else if neighbors(start+2, 500, board) < 8 {
							fmt.Printf("%d %d\n", start+2, 500)
						} else if neighbors(start+2, 501, board) < 8 {
							fmt.Printf("%d %d\n", start+2, 501)
						} else {
							break
						}
					}
				}
			}
		}
	}
}
