package main

import (
	"fmt"
)

type MaxHeap struct {
	heap []int
}

func (mh *MaxHeap) Push(v int) {
	mh.heap = append(mh.heap, v)
	mh.up(len(mh.heap) - 1)
}

func (mh *MaxHeap) Pop() int {
	n := len(mh.heap)
	last := n - 1
	mh.swap(mh.heap, 0, last)
	mh.down(mh.heap, 0, last)

	last = mh.heap[last]
	mh.heap = mh.heap[:n-1]
	return last
}

func (mh *MaxHeap) up(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || mh.more(mh.heap, parent, i) {
			break
		}
		mh.swap(mh.heap, parent, i)
		i = parent
	}
}

func (mh *MaxHeap) Heapify(arr []int) []int {
	n := len(arr)
	lastNonLeafIndex := (len(arr) - 2) / 2 // well known formula to find - last non leaf element
	for lastNonLeafIndex >= 0 {
		mh.down(arr, lastNonLeafIndex, n) //arr, index, size
		lastNonLeafIndex--
	}
	return arr
}

func (mh *MaxHeap) more(arr []int, i, j int) bool {
	return arr[i] > arr[j]
}

func (mh *MaxHeap) down(arr []int, i, n int) {
	parent := i
	for {
		c1 := 2*parent + 1 // first(left) child index
		if c1 >= n {       //if there are no children
			break
		}
		c2 := c1 + 1 // second(right) child index
		swap := c1   // index of child to swap with parent
		if c2 < n && mh.more(arr, c2, c1) {
			swap = c2
		}
		if mh.more(arr, parent, swap) {
			break
		}
		mh.swap(arr, parent, swap)
		parent = swap
	}

}

func (mh *MaxHeap) swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func findKthLargest(nums []int, k int) int {
	mh := &MaxHeap{}
	for _, v := range nums {
		mh.Push(v)
	}
	for i := 0; i < k-1; i++ {
		mh.Pop()
	}
	return mh.Pop()
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
	// sortedNums := mh.Heapify(nums)
	// fmt.Println(sortedNums)

	

}
