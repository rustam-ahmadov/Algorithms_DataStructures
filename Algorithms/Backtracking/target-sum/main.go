package main

import (
	"fmt"
)

// Constraints:

// 1 <= nums.length <= 20
// 0 <= nums[i] <= 1000
// 0 <= sum(nums[i]) <= 1000
// -1000 <= target <= 1000

func findTargetSumWays(nums []int, target int) int {
	return helper(nums, target, 0)
}

func helper(nums []int, target int, cur int) (ways int) {
	if len(nums) == 0 {
		if cur == target {
			return 1
		} else {
			return 0
		}
	}

	plus := helper(nums[1:], target, cur+nums[0])
	minus := helper(nums[1:], target, cur-nums[0])

	return plus + minus
}

func findPrintTargetSumWays(nums []int, target int) int {
	return helperPrint(nums, target, 0, "")
}

func helperPrint(nums []int, target int, cur int, exp string) (ways int) {
	if len(nums) == 0 {
		if cur == target {
			fmt.Println(exp)
			return 1
		} else {
			return 0
		}
	}
	plus := helperPrint(nums[1:], target, cur+nums[0],exp +"+" + string(byte(48+nums[0])))
	minus := helperPrint(nums[1:], target, cur-nums[0],exp +"-" + string(byte(48+nums[0])))

	return plus + minus
}


func findTargetSumWaysMemo(nums []int, target int) int {
	return helper(nums, target, 0)
}

func helperMemo(nums []int, target int, cur int) (ways int) {
	if len(nums) == 0 {
		if cur == target {
			return 1
		} else {
			return 0
		}
	}

	plus := helper(nums[1:], target, cur+nums[0])
	minus := helper(nums[1:], target, cur-nums[0])

	return plus + minus
}

func main() {
	nums := []int{1,1,1,1,1}
	target := 3
	fmt.Println(findPrintTargetSumWays(nums, target))

	str := "s"
	str = str+"a"
	fmt.Println(str)
}
