package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	helper(candidates, target, &res, nil, 0)
	return res
}

func helper(candidates []int, target int, res *[][]int, cur []int, sum int) {
	if sum == target {
		*res = append(*res, cur)
		return
	}
	if sum > target {
		return
	}
	for i, v := range candidates {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if v > target {
			break
		}
		new := append(append([]int{}, cur...), v)
		sum += v
		helper(candidates[i+1:], target, res, new, sum)
		sum -= v
	}
}

func main() {
	c := []int{10, 1, 2, 7, 6, 1, 5}
	t := 8
	fmt.Println(combinationSum2(c, t))
}
