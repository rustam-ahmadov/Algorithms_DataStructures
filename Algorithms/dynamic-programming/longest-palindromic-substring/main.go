package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	max := 0
	var res string
	for i := 0; i < n; i++ {
		even, evenPalindrome := lengthAndPalindromeItself(s, n, i, i+1)
		odd, oddPalindrome := lengthAndPalindromeItself(s, n, i-1, i+1)
		if max < even {
			res = evenPalindrome
			max = even
		}
		if max < odd+1 {
			res = oddPalindrome
			max = odd + 1
		}
	}
	return res
}

func lengthAndPalindromeItself(s string, n, i, j int) (length int, palindrome string) {
	l := 0
	for i >= 0 && j < n {
		if s[i] != s[j] {
			return l, s[i+1 : j]
		}
		l += 2
		i--
		j++
	}
	return l, s[i+1 : j]
}

func main() {
	s := "a"
	fmt.Println(longestPalindrome(s))
}