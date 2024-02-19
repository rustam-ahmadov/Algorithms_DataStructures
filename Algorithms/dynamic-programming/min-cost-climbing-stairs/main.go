package main

import "fmt"

// Constraints:

// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999

//----------------recursion DFS

// var memo [1000]int

// func minCostClimbingStairs(cost []int) int {
// 	one := helper(cost, 0)
// 	two := helper(cost, 1)
// 	memo = [1000]int{}
// 	if one < two {
// 		return one
// 	}
// 	return two
// }

// func helper(cost []int, i int) int {
// 	if i >= len(cost) {
// 		return 0
// 	}

// 	if memo[i] != 0 {
// 		return memo[i]
// 	}

// 	memo[i] = cost[i]
// 	one := helper(cost, i+1)
// 	two := helper(cost, i+2)

// 	if one < two {
// 		memo[i] += one
// 	} else {
// 		memo[i] += two
// 	}
// 	return memo[i]
// }

//-------------------------Iterative

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	first, second := cost[0], cost[1]
	for i := 2; i < n; i++ {
		cur := cost[i] + min(first, second)
		first = second
		second = cur
	}
	return min(first, second)
}

func main() {
	cost0 := []int{1,2,3,1,5}
	fmt.Println(minCostClimbingStairs(cost0))
}
