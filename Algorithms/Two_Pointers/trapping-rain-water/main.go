package main

// Constraints:

// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

func trap(height []int) int {
	l, r := 0, len(height) - 1
	maxL, maxR:= 0,0 
	total:= 0
	for l<r {
		if height[l] <= height[r]{
			if height[l] > maxL{
				maxL = height[l]
			}else {
				total+= maxL - height[l]
			}
			l++
		}else {
			if height[r] > maxR{
				maxR = height[r]
			}else {
				total+= maxR - height[r]
			}
			r--
		}
	}
	return total
}

func main() {
	a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	println(trap(a))
}
