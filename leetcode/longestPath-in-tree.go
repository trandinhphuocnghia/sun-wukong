package leetcode

import (
	"fmt"
	"sunwukong/modal"
)

// post checking
// No matter `LEFT->right->root or Right->left->root`(*) , they are all start from root, go straight to the deepest Node follow the key starting of (*), then follow the (*) go up to root.
func findHeight(node *modal.TreeNode, maxDiameter *int, trackingPath *[]int) (int, []int) {
	if node == nil {
		return 0, []int{}
	}

	// Recursively find the height of left and right subtrees
	leftHeight, leftPath := findHeight(node.Left, maxDiameter, trackingPath)
	rightHeight, rightPath := findHeight(node.Right, maxDiameter, trackingPath)
	fmt.Println("maxDiameter", *maxDiameter)

	fmt.Println("node: ", node.Val)
	fmt.Println("left :", leftHeight)
	if node.Left != nil {
		fmt.Println("Left Node :", node.Left.Val)
	}
	fmt.Println("right", rightHeight)
	if node.Right != nil {
		fmt.Println("Right Node :", node.Right.Val)
	}
	fmt.Println("-----")

	//for the hole tree
	if leftHeight+rightHeight > *maxDiameter {
		*maxDiameter = leftHeight + rightHeight
		*trackingPath = append(leftPath, node.Val)
		*trackingPath = append(*trackingPath, rightPath...)
	}

	//for current node check the under of it
	// Return the height of the current node
	if leftHeight > rightHeight {
		return leftHeight + 1, append(leftPath, node.Val)

	} else {
		return rightHeight + 1, append(rightPath, node.Val)
	}
}

func TrackingTreeTraversal(root *modal.TreeNode) {
	var maxDiameter int
	var trackingPath []int

	findHeight(root, &maxDiameter, &trackingPath)

	fmt.Println(maxDiameter)
	fmt.Println(trackingPath)

}
