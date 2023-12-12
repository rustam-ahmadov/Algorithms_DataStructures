package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	left:= isSameTree(p.Left, q.Left)
	right:= isSameTree(p.Right, q.Right)

	if left == true && right == true {
		return p.Val == q.Val
	}

	return false
}
