package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	arr := []int{root.Val}
	left := rightSideView(root.Left)
	right := rightSideView(root.Right)

	if right != nil && left == nil {
		return append(arr, right...)
	}
	if right == nil && left != nil {
		return append(arr, left...)
	}
	if left != nil {
		if len(right) >= len(left) {
			return append(arr, right...)
		}
		arr = append(arr, right...) // first append right side
		return append(arr, left[len(right):]...)
	}
	return arr
}
