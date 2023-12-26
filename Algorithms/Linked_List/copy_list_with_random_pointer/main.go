package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//my solution with 2 tables 0(n) time and space
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return head
	}

	m1 := make(map[*Node]int)
	m2 := make(map[int]*Node)

	var cpHead *Node = &Node{Val: head.Val}
	m1[head] = 0
	m2[0] = cpHead

	cpTemp := cpHead
	temp := head.Next
	for i := 1; temp != nil; i++ {
		cpTemp.Next = &Node{Val: temp.Val}
		cpTemp = cpTemp.Next

		m1[temp] = i
		m2[i] = cpTemp
		temp = temp.Next
	}

	cpTemp = cpHead
	for head != nil {
		i, isPresent := m1[head.Random]
		if isPresent == false {
			cpTemp.Random = nil
		} else {
			cpTemp.Random = m2[i]
		}
		cpTemp = cpTemp.Next
		head = head.Next
	}
	return cpHead
}

//another solution with 2 cycle and 1 table
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	for temp:= head; temp != nil; temp = temp.Next {
		m[temp] = &Node{Val: temp.Val}
	}

	for temp:= head; temp != nil; temp = temp.Next {
		m[temp].Next = m[temp.Next]
		m[temp].Random = m[temp.Random]
	}
	return m[head]
}

func main() {

	head1 := &Node{Val: 7, Next: nil, Random: nil}
	head2 := &Node{Val: 13, Next: nil, Random: head1}
	head3 := &Node{Val: 11, Next: nil, Random: nil}
	head4 := &Node{Val: 10, Next: nil, Random: nil}
	head5 := &Node{Val: 1, Next: nil, Random: nil}

	head1.Next = head2
	head2.Next = head3
	head3.Next = head4
	head4.Next = head5

	head1.Random = nil
	head2.Random = head1
	head3.Random = head5
	head4.Random = head3
	head5.Random = head1

	copyRandomList(head1)
}
