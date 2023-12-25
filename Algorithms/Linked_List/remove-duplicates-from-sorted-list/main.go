package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//1 1 1 2 3 4 4
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	m := make(map[int]int)
	cur := head
	before := head
	for cur != nil {
		m[cur.Val]++

		if m[cur.Val] == 1 {
			before = cur
		} else if m[cur.Val] > 1 {
			before.Next = cur.Next
		}
		cur = cur.Next
	}
	return head
}
