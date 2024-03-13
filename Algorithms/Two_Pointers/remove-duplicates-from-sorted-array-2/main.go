package main

func removeDuplicates(nums []int) int {
    n := len(nums)
    if n == 0 || n == 1 {
        return n
    }

    l, r := 0, 1
    two := false 
    res := 1
    for r < n {
        if nums[l] == nums[r] && !two {
            l++
            two = true
            res ++
            nums[l] = nums[r]
        }else if nums[l] != nums[r] {
            two = false
            l ++
            nums[l] = nums[r]
            res++
        } 
        r++
    }
    return res
}