func sortColors(nums []int) {
	m := make(map[int]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		m[nums[i]]++
	}

	i := 0
	for m[0] != 0 {
		nums[i] = 0
		m[0]--
		i++
	}
	for m[1] != 0 {
		nums[i] = 1
		m[1]--
		i++
	}
	for m[2] != 0 {
		nums[i] = 2
		m[2]--
		i++
	}
}