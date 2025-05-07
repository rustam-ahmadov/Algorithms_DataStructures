package main

import (
	"fmt"
	"runtime"
)

//Constraints:
//
//1 <= nums.length <= 104
//-104 <= nums[i] <= 104
//nums contains distinct values sorted in ascending order.
//-104 <= target <= 104

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {

		m := (l + r) / 2
		if nums[m] == target {
			return m
		}

		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	if nums[l] < target {
		return l + 1
	}
	return l
}

func searchInsertLinear(nums []int, target int) int {
	for i, v := range nums {
		if target < v {
			return i
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 3}
	target := 0
	res := searchInsert(nums, target)
	fmt.Println(res)

	nums = []int{1, 3, 5, 6, 9}
	target = 10
	res = searchInsert(nums, target)
	fmt.Println(res)

	nums = []int{1, 3, 5, 6, 9}
	target = 5
	res = searchInsert(nums, target)
	fmt.Println(res)

	nums = []int{1, 3, 5, 6, 9, 20, 40, 50, 60, 70, 80, 90, 100, 102, 203, 405, 607, 706}
	target = 608
	res = searchInsert(nums, target)
	fmt.Println(res)

	nums = []int{1, 3, 5, 6, 9, 20, 40, 50, 60, 70, 80, 90, 100, 102, 203, 405, 607, 706}
	target = 2
	res = searchInsert(nums, target)
	fmt.Println(res)

	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumCPU())

}
