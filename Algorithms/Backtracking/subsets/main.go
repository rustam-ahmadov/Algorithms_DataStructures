package main

import "fmt"

// Constraints:

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.

//iterative approach O(2^n)
func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, num := range nums {
		for _, sub := range res {
			new := append([]int{}, sub...) // new allocated copy of sub...
			new = append(new, num)
			res = append(res, new)
		}
	}
	return res
}

func subsetsBack(nums []int) [][]int {
	result := [][]int{}
	// Start backtracking from index 0
	helper(nil, &result, 0, nums)
	return result
}

func helper(currentSubset []int, result *[][]int, start int, nums []int) {
	// Add the current subset to the result
	*result = append(*result, append([]int{}, currentSubset...))

	// Explore all possible subsets
	for i := start; i < len(nums); i++ {
		// Choose
		currentSubset = append(currentSubset, nums[i])

		// Explore
		helper(currentSubset, result, i+1, nums)

		// Unchoose (backtrack)
		currentSubset = currentSubset[:len(currentSubset)-1]
	}
}

func main() {
	nums := []int{1, 2, 3, 4}

	fmt.Println(subsetsBack(nums))
}
