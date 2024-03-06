func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[byte(s[i])]++
	}
	odds := []int{}
	maxOdd := 0
	res := 0
	for _, v := range m {
		if v%2 == 0 {
			res += v
		}else if v % 2 != 0 {
			maxOdd = max(maxOdd, v)
			odds = append(odds, v)
		} 
	}
	res+=maxOdd

	excludeMaxOdd := false
	for _, v := range odds {
		if v == maxOdd && excludeMaxOdd == false {
			excludeMaxOdd = true
			continue
		}else if v > 2 {
			res += v - 1
		}
	}
	return res 
}