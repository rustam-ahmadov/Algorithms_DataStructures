package main

import "fmt"

// Constraints:

// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board and word consists of only lowercase and uppercase English letters.

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if isWord(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func isWord(board [][]byte, y, x int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if x >= 0 && x < len(board[0]) && y >= 0 && y < len(board) && board[y][x] != '%' && board[y][x] == word[0] {
		board[y][x] = '%'
		if isWord(board, y, x-1, word[1:]) || isWord(board, y, x+1, word[1:]) || isWord(board, y-1, x, word[1:]) || isWord(board, y+1, x, word[1:]) {
			return true
		}
		board[y][x] = word[0]
	}
	return false
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED"))
}
