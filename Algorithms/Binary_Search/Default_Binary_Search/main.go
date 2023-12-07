package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	size := 3
	arr := make([]int, size)
	fillArrWRandom(arr, size)
	printArr(arr)

	slices.Sort(arr)
	printArr(arr)

	isPresent := search(arr, size, 1)
	fmt.Printf("%t", isPresent)
}

func fillArrWRandom(arr []int, size int) {

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(5) + 1
	}
}

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func search(arr []int, size, v int) bool {
	l, r := 0, size-1
	
	for l <= r {
		m := (r + l) / 2
		if v == arr[m] {
			return true
		}

		if v < arr[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
