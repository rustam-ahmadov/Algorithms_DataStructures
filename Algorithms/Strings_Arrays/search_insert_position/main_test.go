package main

import (
	"fmt"
	"testing"
)

// Helper function to generate a sorted array of a given size
func generateSortedArray(size int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i * 2 // Each element is 2*i to keep the array sorted
	}
	return nums
}

func BenchmarkSearchInsert(b *testing.B) {

	cases := []struct {
		nums   []int
		target int
		size   int // Добавим размер для отчетов
	}{
		{nums: generateSortedArray(10), target: 11, size: 10},
		{nums: generateSortedArray(100_000), target: 100_001, size: 100_000},
		{nums: generateSortedArray(100_000_000), target: 100_000_001, size: 100_000_000},
	}

	for _, v := range cases {
		b.Run(fmt.Sprintf("binary: Size_%d", v.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				searchInsert(v.nums, v.target)
			}
		})
		b.Run(fmt.Sprintf("linear: Size_%d", v.size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				searchInsertLinear(v.nums, v.target)
			}
		})
	}
}
