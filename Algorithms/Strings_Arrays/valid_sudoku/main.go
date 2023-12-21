package main

import (
	"fmt"
)

// Constraints:
// 		board.length == 9
// 		board{i}.length == 9
// 		board{i}{j} is a digit 1-9 or '.'.

func isValidSudoku(board [][]byte) bool {

	var m map[byte]int = make(map[byte]int)

	//horizontal check
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			sym := board[i][j]
			if sym == '.' {
				continue
			}

			if m[sym]++; m[sym] > 1 {
				return false
			}
		}
		for k := range m {
			delete(m, k)
		}
	}

	//vertical check
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {

			sym := board[i][j]
			if sym == '.' {
				continue
			}

			if m[sym]++; m[sym] > 1 {
				return false
			}
		}
		for k := range m {
			delete(m, k)
		}

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {

				for i_i := i * 3; i_i < i*3+3; i_i++ {

					for j_j := j * 3; j_j < j*3+3; j_j++ {

						sym := board[i_i][j_j]
						if sym == '.' {
							continue
						}

						if m[sym]++; m[sym] > 1 {
							return false
						}
					}
				}

				for k := range m {
					delete(m, k)
				}
			}
		}
	}

	return true
}

func main() {

	board := [][]byte{
		{'5', '3', '4', '.', '7', '.', '.', '.', '.'},
		{'6', '4', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	isValid := isValidSudoku(board)
	fmt.Print(isValid)
}
