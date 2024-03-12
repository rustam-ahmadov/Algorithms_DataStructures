package main

func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	res, cur := 0, 0
	l, r := 0, 0
	m := make(map[int]int)
	for r < n { 
		num := nums[r]
		m[num]++
        cur += num
		for m[num] > 1 {
			numToErase := nums[l]
			m[numToErase]--
			cur -= numToErase
			l++
		}
        res = max(cur, res)
        r++
	}
    return res
}

func main(){
	
}