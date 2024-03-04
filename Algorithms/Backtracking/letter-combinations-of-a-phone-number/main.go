package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := make(map[byte]string)
	m[byte('2')] = "abc"
	m[byte('3')] = "def"
	m[byte('4')] = "ghi"
	m[byte('5')] = "jkl"
	m[byte('6')] = "mno"
	m[byte('7')] = "pqrs"
	m[byte('8')] = "tuv"
	m[byte('9')] = "wxyz"

	res := []string{}
	helper(&res, digits, "", m)
	return res
}

func helper(res *[]string, digits, cur string, m map[byte]string) {
	if len(digits) == 0 {
		*res = append(*res, cur)
		return
	}

	for _, v := range m[byte(digits[0])] { // v = "a"
		helper(res, digits[1:], cur+string(v), m)
	}
}
