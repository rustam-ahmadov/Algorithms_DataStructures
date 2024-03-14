package main

func longestSubarray(nums []int) int {
	n := len(nums)
	res := 0
	zeros := 0
	for l, r := 0, 0; r < n; r++ {
		if nums[r] == 0 {
			zeros++
		}
		if zeros > 1 {
			for nums[l] != 0 {
				l++
			}
			zeros--
			l++
		}
		cur := r - l + 1 - zeros
		res = max(cur, res)
	}
	if res == n {
		return res - 1
	}
	return res
}
