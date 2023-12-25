package main

// Constraints:
// 	m == s.length
// 	n == t.length
// 	1 <= m, n <= 105
// 	s and t consist of uppercase and lowercase English letters.

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)

	if sLen < tLen { //substring is bigger
		return ""
	}
	//insert
	m := make(map[byte]int)
	for i := 0; i < tLen; i++ {
		m[t[i]]++
	}

	l, r := 0, 0
	counter := tLen
	startIndex, minLen := 0, 0

	for r < sLen {
		curSym := s[r]
		if m[curSym] > 0 {
			counter--
		}
		//If char does not exist in t, m[curSym] will be negative.
		m[curSym]--
		r++
		// When we found a valid window, move start to find smaller window.
		for counter == 0 {
			curLen := r - l
			if curLen < minLen || minLen == 0 {
				startIndex = l
				minLen = curLen
			}

			curSym := s[l]
			m[curSym]++
			// When char exists in t, increase counter.
			if m[curSym] > 0 {
				counter++
			}
			l++
		}
	}
	if minLen != 2147483647 {
		return s[startIndex : startIndex+minLen]
	}
	return ""
}

func main() {
	minWindow("hello", "ell")
}
