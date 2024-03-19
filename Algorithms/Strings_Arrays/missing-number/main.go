package main

func missingNumber(nums []int) int {
	n := len(nums)
	res := 0
	for j,i := 0, 1; i <= n ; i++ {
		res += i
		res -= nums[j]
        j++
	}
	return res
}