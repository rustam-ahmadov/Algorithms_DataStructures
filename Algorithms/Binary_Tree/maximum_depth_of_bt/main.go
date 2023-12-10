package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)

	max := 1
	if leftMax < rightMax {
		max += rightMax
	}else {
		max += leftMax
	}
	return max
}


func main() {
}
