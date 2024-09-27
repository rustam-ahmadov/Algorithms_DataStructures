package main

import "fmt"

//Constraints:
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 104

// O(n ^ 2) cause it returns both pointers at the position of l pointer
func minSubArrayLenOn2(target int, nums []int) int {
	n := len(nums)

	res, sum, sub := 0, 0, 0
	l, r := 0, 0
	for r < n {
		sum += nums[r]
		if sum >= target {
			sub = (r - l) + 1
			if res == 0 || res > sub {
				res = sub
			}
			l, r = l+1, l
			sum = 0
		}
		r++
	}
	return res
}

func minSubArrayLenOn(target int, nums []int) int {
	n := len(nums)
	l, r := 0, 0
	res, sum, sub := 0, 0, 0
	for r < n {
		sum += nums[r]
		for sum >= target {
			sub = (r - l) + 1
			if res == 0 || res > sub {
				res = sub
			}
			sum -= nums[l]
			l++
		}
		r++
	}
	return res
}

func main() {
	target1 := 4
	nums1 := []int{1, 4, 4}
	res := minSubArrayLenOn(target1, nums1)
	fmt.Println(res)

	target2 := 7
	nums2 := []int{2, 3, 1, 2, 4, 3}
	res = minSubArrayLenOn(target2, nums2)
	fmt.Println(res)
}
