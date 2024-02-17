package main

import "fmt"

// Constraints:

// 1 <= n <= 45

var memo [45]int

func climbStairs(n int) int {

	if n <= 1 {
		return n
	}
	if n == 2 {
		return 2
	}
	if memo[n-1] != 0 {
		return memo[n-1]
	}
	memo[n-1] = climbStairs(n-1) + climbStairs(n-2)
	return memo[n-1]
}

func climbStairsDyp(n int) int {
	if n == 1 {
		return 1
	}
	f, s := 1, 1 // s = res of prev
	res := f + s
	for i := 2; i < n; i++ {
		f = s
		s = res
		res = f + s
	}
	return res
}

func main() {
	fmt.Println(climbStairsDyp(4))
}
