package main

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

func longestConsecutive(nums []int) int {
    m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	longest := 0
	for k, _ := range m {
		delete(m, k)
        cur := 1
		l, r := k - 1, k + 1
        for m[l] {
            delete(m, l)
            l--
            cur++
        }
        for m[r] {
            delete(m, r)
            r++
            cur++
        }
		if cur > longest {
			longest = cur
		}
	}
	return longest
}
