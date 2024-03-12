package main

func sortedSquares(nums []int) []int {
    n := len(nums)
    res := make([]int, n)
    i := n - 1
    l, r := 0, n - 1
    for l <= r {
        if nums[l] * nums[l] < nums[r] * nums[r] {
            res[i] = nums[r] * nums[r]
            r--
        }else {
            res[i] = nums[l] * nums[l]
            l++
        }
        i--
    }
    return res
}