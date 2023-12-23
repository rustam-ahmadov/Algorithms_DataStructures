package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Constraints:
// 	The number of nodes in the list is sz.
// 	1 <= sz <= 30
// 	0 <= Node.val <= 100
// 	1 <= n <= sz

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	f, s := head, head

	//fast till n from start 1, 2, 3, 4, 5, 6  and n = 2 f will be at (2)
	for startNode := 1; startNode < n; startNode++ {
		f = f.Next
	}

	if s == f && f.Next == nil { // we have only one node in listNode
		return nil
	}

	if s != f && f.Next == nil { // f on the end that means that we need to remove our first node
		return s.Next
	}

	//fast till the end and slow till the end - n, f on 6 slow on 4, prev on 3
	var prev *ListNode = s
	for f.Next != nil {
		prev = s
		s = s.Next
		f = f.Next
	}

	prev.Next = s.Next
	return head
}
