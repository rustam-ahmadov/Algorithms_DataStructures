package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	depthLeft, diameterLeft := depthDiameter(root.Left)
	depthRight, diameterRight := depthDiameter(root.Right)

	currDiameter := depthLeft + depthRight
	maxDiameter := max(diameterLeft, diameterRight)
	return max(currDiameter, maxDiameter)
}

func depthDiameter(root *TreeNode) (depth, diameter int) {
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 1, 0
	}
	depthLeft, diameterLeft := depthDiameter(root.Left)
	depthRight, diameterRight := depthDiameter(root.Right)

	currDepth := max(depthLeft, depthRight) + 1
	currDiameter := depthLeft + depthRight
	childMaxDiameter := max(diameterLeft, diameterRight)

	if currDiameter > childMaxDiameter {
		return currDepth, currDiameter
	}
	return currDepth, childMaxDiameter
}

func main() {
	root :=
		&TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: nil,
		}
	diameterOfBinaryTree(root)
}
