package main

import (
	"fmt"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	max := slices.Max(piles)
	if len(piles) == h { //optimization
		return max
	}

	l, r := piles[0], max
	min := 0
	for l <= r {
		m := (l + r) / 2

		hours := 0
		for i := 0; i < len(piles); i++ {
			hours += piles[i] / m
			if piles[i]%m > 0 {
				hours++
			}
			if hours > h {
				break
			}
		}

		if hours <= h {
			min = m
		}

		if hours > h {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return min
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	h := 6
	res := minEatingSpeed(piles, h)
	fmt.Printf("min eating speed: %d\n", res)
}
