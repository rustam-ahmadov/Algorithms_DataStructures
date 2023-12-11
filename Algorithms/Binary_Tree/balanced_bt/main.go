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
	
	depthLeft := depth(root.Left)
	depthRight := depth(root.Right)

	if depthLeft == depthRight {
		return true
	}
	if depthLeft == depthRight + 1 {
		return true
	} 
	if depthRight == depthLeft + 1 {
		return true
	}
	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	depthLeft := depth(root.Left)
	depthRight := depth(root.Right)

	return max(depthLeft, depthRight) + 1
}

func main() {

}
