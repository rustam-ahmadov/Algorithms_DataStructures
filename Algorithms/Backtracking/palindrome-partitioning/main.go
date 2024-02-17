package main

import "fmt"

// Constraints:

// 1 <= s.length <= 16
// s contains only lowercase English letters.

func partition(s string) [][]string {

	res := [][]string{}
	helper(&res, s, []string{})
	return res
}

func helper(res *[][]string, s string, cur []string) {
	if len(s) == 0 {
		*res = append(*res, append([]string{}, cur...))
		return
	}

	var sub string
	for i := 0; i < len(s); i++ {
		sub = sub + string(s[i])
		if !isPalindrome(sub) {
			continue
		}
		newCur := append(append([]string{}, cur...), sub)
		helper(res, s[i+1:], newCur)
	}
}

func isPalindrome(substring string) bool {
	if len(substring) == 0 {
		return false
	}
	for i := 0; i < len(substring)/2; i++ {
		l := substring[i]
		r := substring[len(substring)-i-1]
		if l != r {
			return false
		}
	}
	return true
}

func main() {

	s := "aab" // size = 8
	fmt.Println(isPalindrome(s))

	fmt.Println(partition(s))
}
