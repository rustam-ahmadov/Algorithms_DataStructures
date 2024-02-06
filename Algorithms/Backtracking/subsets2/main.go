package main

import "fmt"

// Constraints:

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	
	// helper(&res, nums, nil, 0, 1)
	// return res

	for i := 0; i < len(nums); i++ {
		cur := []int{nums[i]}
		res = append(res, cur)
		for j := i + 1; j < len(nums); j++ {
			cur = append(append([]int{}, cur...), nums[j])
			res = append(res, cur)
		}
	}
	return res
}

// func helper(res *[][]int, nums []int, cur []int, l int, r int) {
// 	for i := l; i < r && i < len(nums); i++ {
// 		*res = append(*res, []int{nums[i]})
// 		new := append(append([]int{}, cur...), nums[i])
// 		helper(res, nums, new, l+1, l+2)
// 	}
// 	*res = append(*res, cur)
// }

func main() {
	nums := []int{1, 2, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
