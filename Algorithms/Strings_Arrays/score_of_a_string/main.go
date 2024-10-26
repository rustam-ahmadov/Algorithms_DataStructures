package main

import "fmt"

//Constraints:
//
//2 <= s.length <= 100
//s consists only of lowercase English letters.

func scoreOfString(s string) int {

	n := len(s)
	res := 0

	for i := 0; i < n-1; i++ {
		l := int(s[i])
		r := int(s[i+1])

		lr := l - r
		if lr < 0 {
			lr *= -1
		}
		res += lr
	}
	return res
}

func main() {
	str1 := "hello"
	score := scoreOfString(str1)
	fmt.Println(score)

	str2 := "əустам"
	score = scoreOfString(str2)
	fmt.Println(score)

	str3 := "ab"
	score = scoreOfString(str3)
	fmt.Println(score)

	str4 := "abc"
	score = scoreOfString(str4)
	fmt.Println(score)
}
