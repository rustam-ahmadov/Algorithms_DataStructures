func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[byte(s[i])]++
	}
	one, oddMoreThanOne := false, false
	res := 0
	for _, v := range m {
		if v%2 == 0 {
			res += v
		} else if v%2 != 0 && v != 1 {
			oddMoreThanOne = true
			res += v - 1
		} else {
			one = true
		}
	}
	if !one && !oddMoreThanOne {
		return res
	}
	return res + 1
}