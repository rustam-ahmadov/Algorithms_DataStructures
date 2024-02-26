package main

import "fmt"

// Constraints:

// 	1 <= nums.length <= 100
// 	0 <= nums[i] <= 1000

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	return max(helper(nums[:len(nums)-1]), helper(nums[1:]))
}

func helper(nums []int) int {
	prev, cur := 0,0
	for i := 0; i < len(nums); i++ {
		temp := cur
		cur = max(nums[i]+prev, cur)
		prev = temp
	}
	return cur
}

func main() {
	nums := []int{2, 3, 2}
	// nums:= []int{2,6,2}
	// nums:= []int{5,2,1,4}

	fmt.Println(rob(nums))
}
