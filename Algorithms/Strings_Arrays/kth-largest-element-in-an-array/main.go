package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap struct {
	data []int
}

func (mh *MaxHeap) Len() int {
	return len(mh.data)
}

func (mh *MaxHeap) Less(i, j int) bool {
	return mh.data[i] > mh.data[j]
}

func (mh *MaxHeap) Swap(i, j int) {
	mh.data[i], mh.data[j] = mh.data[j], mh.data[i]
}

func (mh *MaxHeap) Push(x interface{}) {
	mh.data = append(mh.data, x.(int))
}

func (mh *MaxHeap) Pop() interface{} {
	n := len(mh.data)
	last := mh.data[n-1]
	mh.data = mh.data[:n-1]
	return last
}

func findKthLargest(nums []int, k int) int {
	mh := &MaxHeap{}
	heap.Init(mh)
	for _, v := range nums {
		heap.Push(mh, v)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(mh)
	}
	return heap.Pop(mh).(int)
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println(findKthLargest(nums, 4))
}
