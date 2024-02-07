package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Constraints:

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}

	// O(n (log n))
	sort.SliceStable(nums, func(i, j int) bool { return nums[i] < nums[j] })
	helper(&res, nums, nil, 0)
	return res
}

func helper(res *[][]int, nums []int, cur []int, index int) {
	*res = append(*res, cur)
	prev := []int{}
	for i := index; i < len(nums); i++ {
		new := append(append([]int{}, cur...), nums[i])
		if !reflect.DeepEqual(prev, new) {
			helper(res, nums, new, i + 1)
		}
		prev = new
	}
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
