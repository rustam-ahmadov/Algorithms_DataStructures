package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
		return root
	}
	invert(root)
	return root
}

func invert(node *TreeNode){
    if node.Left != nil {
        invert(node.Left)
    }
    if node.Right != nil {
        invert(node.Right)
    }
    temp := node.Right
	node.Right = node.Left
	node.Left = temp
}

func main() {
	
	root := &TreeNode{4,
		&TreeNode{2,
			&TreeNode{Val: 1},
			&TreeNode{Val: 3},
		},
		&TreeNode{7,
			&TreeNode{Val: 6},
			&TreeNode{Val: 9},
		},
	}

	invertTree(root)
}
