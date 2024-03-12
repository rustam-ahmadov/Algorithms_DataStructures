package main

import (
	"fmt"
	// "strings"
)

//O(n ^ 3)
// func countSubstrings(s string) int {
// 	res := 0
// 	var sb strings.Builder
// 	for i := 0; i < len(s); i++ {
// 		sb.WriteString(string(s[i]))
// 		if isPalindrome(sb.String()) {
// 			res += 1
// 		}
// 		for j := i + 1; j < len(s); j++ {
// 			sb.WriteString(string(s[j]))
// 			if isPalindrome(sb.String()) {
// 				res += 1
// 			}
// 		}
// 		sb.Reset()
// 	}
// 	return res
// }

func isPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

// 0(n ^ 2)
func countSubstrings(s string) int {
	res := 0
	n := len(s)
	for i := 0; i < n; i++ {
		odd := palindromeCount(s, i-1, i+1, n)
		even := palindromeCount(s, i, i+1, n)
		res += odd + even + 1
	}
	return res
}

func palindromeCount(s string, l, r, n int) int {
	count := 0
	for l >= 0 && r < n {
		if s[l] != s[r] {
			return count
		}
		count++
		l--
		r++
	}
	return count
}
func main() {
	s := "aaa"
	fmt.Println(countSubstrings(s))
}
