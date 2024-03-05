package main

import "strings"

func minRemoveToMakeValid(s string) string {
	stack := []int{}
	var prev byte
	for i, v := range s {
		sym := byte(v)
		if sym == '(' || sym == ')' && prev != '(' {
			stack = append(stack, i) //push to stack
			prev = sym
		} else if sym == ')' && prev == '(' {
			stack = stack[:len(stack)-1] //pop from stack
			if len(stack) == 0 {
				prev = byte(' ')
			} else {
				prev = byte(s[stack[len(stack)-1]])
			}
		}
	}
	var res strings.Builder
	for i, j := 0, 0; i < len(s); i++ {
		if j < len(stack) && i == stack[j] {
			j++
			continue
		}
		res.WriteString(string(s[i]))
	}
	return res.String()
}