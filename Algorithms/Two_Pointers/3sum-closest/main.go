package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	res := 0
	diff := math.MaxInt
	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			temp := nums[i] + nums[left] + nums[right]

			curDiff := 0

			if temp < target {
				curDiff = target - temp
				left++
			} else if temp > target {
				curDiff = temp - target
				right--
			} else {
				return temp
			}

			if curDiff < diff {
				diff = curDiff
				res = temp
			}
		}
	}
	return res
}

func main() {
	nums := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	target := -2

	res := threeSumClosest(nums, target)

	fmt.Println(res)
}
