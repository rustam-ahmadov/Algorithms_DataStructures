package main

import (
	"fmt"
)

var closingToOpening = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	str := []byte(s)
	stack := make([]byte, 0)

	for _, v := range str {
		opening, isClosing := closingToOpening[v]
		if !isClosing {
			stack = append(stack, v)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if stack[len(stack)-1] != opening {
			return false
		}

		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func main() {

	// var s1 string = "()"
	var s2 string = "()[]{}"
	// var s3 string = "(]"
	// var s4 string = "({})[]"

	var valid bool = isValid(s2)
	fmt.Printf("string: %s. isValid: %t\n", s2, valid)
}
