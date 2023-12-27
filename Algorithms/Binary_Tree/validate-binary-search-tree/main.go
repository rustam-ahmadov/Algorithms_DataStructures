package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	l := helper(root.Left, math.MinInt64, root.Val)
	r := helper(root.Right, root.Val, math.MaxInt64)

	return l && r
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	l := helper(root.Left, min, root.Val)
	r := helper(root.Right, root.Val, max)

	if !l || !r {
		return false
	}
	return root.Val > min && root.Val < max
}
