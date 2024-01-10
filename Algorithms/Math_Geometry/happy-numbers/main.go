package main

import "fmt"

func isHappy(n int) bool {
	digits := numToDigits(n)
	sum := squareSum(digits)
	m := make(map[int]bool)
	for sum != 1 {
		digits = numToDigits(sum)
		sum = squareSum(digits)
		if m[sum] {
			return false
		}
		m[sum] = true
	}
	return true
}

func squareSum(digits []int) int {
	sum := 0
	for _, v := range digits {
		sum += v * v
	}
	return sum
}

func numToDigits(num int) []int {

	digits := make([]int, 0)
	for num != 0 {
		d := num % 10
		num /= 10
		digits = append(digits, d)
	}
	return digits
}

func main() {
	num := 2
	fmt.Println(isHappy(num))
}
