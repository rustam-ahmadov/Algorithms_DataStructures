package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Constraints:

// k == lists.length
// 0 <= k <= 104
// 0 <= lists[i].length <= 500
// -104 <= lists[i][j] <= 104
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 104.

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
		return nil
	}
	sorted := make([]int, 0)
	for i := 0; i < len(lists); i++ {
		for lists[i] != nil {
			sorted = append(sorted, lists[i].Val)
			lists[i] = lists[i].Next
		}
	}

	if len(sorted) == 0 {
		return nil
	}
	sort.Ints(sorted)
	mergedList := &ListNode{sorted[0], nil}
	cur := mergedList
	for _, v := range sorted[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return mergedList
}

func main() {
	// [[1,4,5],[1,3,4],[2,6]]
	node8 := &ListNode{Val: 5}
	node5 := &ListNode{2, node8}
	node2 := &ListNode{0, node5}

	node7 := &ListNode{Val: 5}
	node4 := &ListNode{4, node7}
	node1 := &ListNode{1, node4}

	node6 := &ListNode{Val: 6}
	node3 := &ListNode{2, node6}

	list := &[]*ListNode{node1, node3, node2}

	merged := mergeKLists(*list)
	fmt.Println(merged)
}
