package main

import "fmt"

var cache [101][101]int = [101][101]int{{1}}

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		cache[m][n] = 1
	}
	if cache[m][n] != 0 {
		return cache[m][n]
	}
	if m > 1 {
		cache[m-1][n] = uniquePaths(m-1, n)

	}
	if n > 1 {
		cache[m][n-1] = uniquePaths(m, n-1)
	}
	return cache[m-1][n] + cache[m][n-1]
}
func main() {
	m, n := 3, 7
	fmt.Println(uniquePaths(m, n))
}
