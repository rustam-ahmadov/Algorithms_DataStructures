package main

import (
	"fmt"
	"sort"
)

// Constraints:

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}

	// O(n (log n))
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] < nums[j] })
	helper(&res, nums, nil)
	return res
}

func helper(res *[][]int, nums []int, cur []int) {
	*res = append(*res, cur)
	for i := 0; i < len(nums); i++ {
		new := append(append([]int{}, cur...), nums[i])
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		helper(res, nums[i+1:], new)
	}
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
