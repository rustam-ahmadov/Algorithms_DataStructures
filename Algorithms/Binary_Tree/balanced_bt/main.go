package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	depthLeft, isBalancedLeft := depthIsBalanced(root.Left)
	depthRight, isBalancedRight := depthIsBalanced(root.Right)

	if isBalancedLeft == false || isBalancedRight == false {
		return false
	}

	if depthLeft == depthRight || depthLeft == depthRight+1 || depthLeft == depthRight-1 {
		return true
	}

	return false
}

func depthIsBalanced(root *TreeNode) (depth int, isBalanced bool) {
	if root == nil {
		return 0, true
	}
	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	depthLeft, isBalancedLeft := depthIsBalanced(root.Left)
	depthRight, isBalancedRight := depthIsBalanced(root.Right)

	if isBalancedLeft == false || isBalancedRight == false {
		return 0, false
	}

	if depthLeft == depthRight || depthLeft == depthRight+1 || depthLeft == depthRight-1 {
		return max(depthLeft, depthRight) + 1, true
	}
	
	return 0, false
}

func main() {

}
