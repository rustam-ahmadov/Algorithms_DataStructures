package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) print() {
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
}

func reorderList(head *ListNode) {

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	secondList := slow
	reversedSecondList := reverse(secondList)
	a := head
	b := reversedSecondList
	for b.Next != nil {
		aNext := a.Next
		bNext := b.Next
		a.Next = b
		b.Next = aNext
		a = aNext
		b = bNext
	}
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversed := head
	cur := head.Next
	prev := head

	for cur != nil {
		next := cur.Next 
		cur.Next = prev  
		prev = cur 
		cur = next  
	}
	reversed.Next = nil 
	return prev
}

func main() {
	list := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4, &ListNode{
						5,nil,
					},
					
				},
			},
		},
	}

	reorderList(list)
	list.print()
}
