package main

// Constraints:

// 	2 <= nums.length <= 105
// 	1 <= nums[i] <= 106

func minOperations(nums []int) int {
	//all nums to hash table
	var m map[int]int = make(map[int]int, 0)
	for _, v := range nums { //i = index, v = value
		m[v]++
	}

	minOp := 0
	for _, v := range m { //k = key, v = value
		if v == 1 {
			return -1
		}

		if v%3 == 0 {
			minOp += v / 3
		} else {
			minOp += v/3 + 1
		}
	}
	return minOp
}

func main() {
	nums := []int{19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19}
	minOperations(nums)
}
