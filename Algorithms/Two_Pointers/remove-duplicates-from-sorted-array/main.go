package main

func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 0 || n == 1 {
        return n
    }
    l, r := 0, 1
    res := 1
    for r < n{
        if nums[l] != nums[r] {
            res ++ 
            l++ 
            nums[l] = nums[r]   
        }
        r++
    }
    return res
}