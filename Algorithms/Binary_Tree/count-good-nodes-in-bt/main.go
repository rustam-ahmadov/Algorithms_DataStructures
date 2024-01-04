package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Constraints:
// 	The number of nodes in the binary tree is in the range [1, 10^5].
// 	Each node's value is between [-10^4, 10^4].

func goodNodes(root *TreeNode) int {
	left := helper(root.Left, root.Val)
	right := helper(root.Right, root.Val) //0

	return left + right + 1
}

func helper(root *TreeNode, prev int) int {
	if root == nil {
		return 0
	}

	p := prev 
	if prev < root.Val {
		p = root.Val 
	} 
	left := helper(root.Left, p)   //4 0
	right := helper(root.Right, p) //0

	res := left + right
	if  prev <= root.Val {
		res += 1
	}
	return res
}

func main() {
	node := &TreeNode{
		3,
		&TreeNode{
			3,
			&TreeNode{
				Val: 4,
			},
			&TreeNode{
				Val: 2,
			},
		},
		nil,
	}

	fmt.Print(goodNodes(node))
}
