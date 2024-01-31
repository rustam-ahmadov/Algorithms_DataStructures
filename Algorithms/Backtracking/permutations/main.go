package main

import "fmt"

// Constraints:

// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.

func permute(nums []int) [][]int {

	res := [][]int{}
	if len(nums) == 1 {
		res = append(res, []int{nums[0]})
	}

	for i := 0; i < len(nums); i++ {
		arrWOIndex := append([]int{}, nums[:i]...)
		arrWOIndex = append(arrWOIndex, nums[i+1:]...)
		perms := permute(arrWOIndex)

		for _, perm := range perms {
			perm = append([]int{nums[i]}, perm...)
			res = append(res, perm)
		}
	}

	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))

}
