package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	prev := nums[0]
	cur := max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		temp := cur
		cur = max(nums[i]+prev, cur)
		prev = temp
	}
	return cur
}

func main() {
	nums := []int{4,1,1, 5}
	fmt.Println(rob(nums))
}
