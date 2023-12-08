/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//recursive
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}

	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var temp *TreeNode = root
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	for temp != nil {
		if val < temp.Val {
			if temp.Left == nil {
				temp.Left = &TreeNode{val, nil, nil}
				break
			}
			temp = temp.Left
		}

		if val > temp.Val {
			if temp.Right == nil {
				temp.Right = &TreeNode{val, nil, nil}
				break
			}
			temp = temp.Right
		}
	}
	return root
}