package main

import "fmt"

// Constraints:

// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sub := make([]int, 0)

	helper(&res, sub, target, candidates, 0)
	return res
}

func helper(res *[][]int, sub []int, target int, candidates []int, index int) {
	sum := 0
	for _, num := range sub {
		sum += num
	}
	if sum == target {
		*res = append(*res, append([]int{}, sub...))
		return
	}
	if sum > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		sub = append(sub, candidates[i])
		helper(res, sub, target, candidates, i)
		sub = sub[:len(sub)-1]
	}
}

func main() {
	input := []int{2, 3, 6, 7}
	target := 7

	res := combinationSum(input, target)
	fmt.Println(res)
}
