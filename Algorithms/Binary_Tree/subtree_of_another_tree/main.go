package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot){
		return true
	}

	if root.Val != subRoot.Val {
		return false				
	}	
	if !isSame(root.Left, subRoot.Left) || !isSame(root.Right, subRoot.Right) {
		return false 
	}
	return true
}

func isSame(p *TreeNode, q *TreeNode ) bool{
	if p == nil || q == nil {
		return p == q
	}
	if !isSame(p.Left, q.Left)|| !isSame(p.Right, q.Right){
		return false
	}
	return p.Val == q.Val
}
