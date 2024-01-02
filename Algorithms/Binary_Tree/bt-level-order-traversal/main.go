package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Constraints:
// 		The number of nodes in the tree is in the range [0, 2000].
// 		-1000 <= Node.val <= 1000

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	matrix := [][]int{{root.Val}}
	left := levelOrder(root.Left)
	right := levelOrder(root.Right)

	lLevel := len(left)
	rLevel := len(right)
	for i := 0; i < lLevel || i < rLevel; i++ {
		if i < lLevel && i < rLevel {
			arr := left[i]
			arr = append(arr, right[i]...)
			matrix = append(matrix, arr)
		} else if i < lLevel {
			matrix = append(matrix, left[i])
		} else {
			matrix = append(matrix, right[i])
		}
	}
	return matrix
}

func main() {
	// head := TreeNode{Val: 1}
	// head.Left = &TreeNode{Val: 2}
	// head.Right = &TreeNode{Val: 3}

	// head.Left.Left = &TreeNode{Val: 4}
	// head.Left.Right = &TreeNode{Val: 5}

	// levelOrder(&head)

	// arr := make([]int, 0)
	// // arr1 := []int{1, 2, 3, 4, 5}
	// arr2 := []int{6, 7, 8, 9, 0}
	// // copy(arr, arr1)
	// // copy(arr[5:], arr2)
	// arr = append(arr, arr2...)
	// fmt.Println(arr)

	matrix := [][]int{{1}}
	arr1 := []int{2, 3}
	arr2 := []int{4, 5}
	arr3 := []int{5, 6}

	matrix = append(matrix, arr1)

	arr4 := arr2
	arr4 = append(arr4, arr3...)
	matrix = append(matrix, arr4)
	fmt.Println(matrix)
}
