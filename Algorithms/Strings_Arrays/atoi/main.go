package main

import "fmt"

func myAtoi(s string) int {

	i := 0
	for i < len(s) {
		if s[i] != ' ' {
			break
		}
		i++
	}

	minus := false
	for i < len(s) {
		if s[i] == '-' {
			minus = true
			i++
		} else if s[i] == '+' {
			i++
		}
		break
	}

	res := 0
	for ; i < len(s); i++ {

		digit := int(s[i]) - 48
		if digit < 0 || digit > 9 { //short path if first is not number (. and a ... z)
			break
		}

		if minus {
			digit *= -1
		}
		res = res*10 + digit

		if res > 2147483647 {
			return 2147483647
		}
		if res < -2147483648 {
			return -2147483648
		}
	}

	return res
}

func main() {
	fmt.Println(myAtoi("456"))
	fmt.Println(myAtoi("00001"))
	fmt.Println(myAtoi("--1"))
	fmt.Println(myAtoi("1337c0d3"))
	fmt.Println(myAtoi("0-1"))
	fmt.Println(myAtoi(" -042"))
	fmt.Println(myAtoi(" --042"))
	fmt.Println(myAtoi("-91283472332"))        //91_283_472_332
	fmt.Println(myAtoi("9223372036854775808")) //91_283_472_332

}
