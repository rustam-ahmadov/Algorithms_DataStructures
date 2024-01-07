package main

import ()

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Constraints:

// 	The number of nodes in the tree is n.
// 	1 <= k <= n <= 104
// 	0 <= Node.val <= 104

// 	Follow up: If the BST is modified often (i.e., we can do insert and delete operations)
//  and you need to find the kth smallest frequently, how would you optimize?

// DFS: Inorder traversal - Time O(n) Space0(n)
//	We don't need to sort this slice cause this is already binary search tree, left is always less than right.
//	First we add left side < then we add middle, root's value < then the right side
func kthSmallest(root *TreeNode, k int) int {
	left := helper(root.Left)
	right := helper(root.Right)
	nums := append(append(left, root.Val), right...)
	return nums[k-1]
}

func helper(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	left := helper(root.Left)
	right := helper(root.Right)
	return append(append(left, root.Val), right...)
}

// Iterative Time O(n) Space is also O(n) but less than in dfs cause we don't create stack frame of
// functions(recursion is also require additional space in stack)
func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]   // previous, cause now root is nil
		stack = stack[:len(stack)-1] // stack = stack - 1 (without current)

		k -= 1
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}

func main() {
	//[3,1,4,null,2]
	n3 := &TreeNode{Val: 4}
	n4 := &TreeNode{Val: 2}
	n2 := &TreeNode{Val: 1, Right: n4}

	n1 := &TreeNode{Val: 3,
		Left:  n2,
		Right: n3,
	}
	kthSmallest(n1, 1)
	// n5:= TreeNode{Val: 3}

}
