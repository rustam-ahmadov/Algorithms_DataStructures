package main

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	helper(&res, []int{}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, k, n)
	return res
}

func helper(res *[][]int, cur []int, nums []int, sum, k, n int) {
	if sum == n && len(cur) == k {
		*res = append(*res, cur)
		return
	}

	if len(cur) == k {
		return
	}

	for i := 0; i < len(nums); i++ {
		new := append([]int{}, cur...)
		new = append(new, nums[i])
		helper(res, new, nums[i+1:], sum+nums[i], k, n)
	}
}


func main() {

}
