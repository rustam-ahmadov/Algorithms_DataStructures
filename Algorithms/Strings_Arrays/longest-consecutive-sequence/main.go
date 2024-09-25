package main

import "fmt"

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109

//func longestConsecutive(nums []int) int {
//    m := make(map[int]bool)
//	for _, v := range nums {
//		m[v] = true
//	}
//
//	longest := 0
//	for k, _ := range m {
//		delete(m, k)
//        cur := 1
//		l, r := k - 1, k + 1
//        for m[l] {
//            delete(m, l)
//            l--
//            cur++
//        }
//        for m[r] {
//            delete(m, r)
//            r++
//            cur++
//        }
//		if cur > longest {
//			longest = cur
//		}
//	}
//	return longest
//}

func longestConsecutiveSequence(nums []int) int {
	m := make(map[int]bool)

	for _, v := range nums {
		m[v] = true
	}

	res := 0
	cur := 0
	for v := range m {
		cur++

		lower := v - 1
		for m[lower] { // lower is exist
			delete(m, lower)
			cur++
			lower--
		}

		greater := v + 1
		for m[greater] { // greater is present
			delete(m, greater)
			cur++
			greater++
		}

		if res < cur {
			res = cur
		}
		cur = 0

		delete(m, v)
	}
	return res
}

func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1} // 100, 200, 4, 1, 3, 2   =  res := 1, 2, 3, 4   constraint = O(n)

	//sort O(n log n)    1, 2, 3, 4 , 100 ,200
	fmt.Println(longestConsecutiveSequence(nums))
}
