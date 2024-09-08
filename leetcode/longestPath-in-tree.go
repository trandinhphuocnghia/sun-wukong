package leetcode

import (
	"fmt"
	"sunwukong/modal"
)

// post checking
// No matter `LEFT->right->root or Right->left->root`(*) , they are all start from root, go straight to the deepest Node follow the key starting of (*), then follow the (*) go up to root.
func postOrderTraversal(node *modal.TreeNode, walkThru *[]modal.WalkThru, theTraversal *[]int) {
	if node == nil {
		return
	}
	var step modal.WalkThru

	//fmt.Println("FROM: ", node.Val)

	if node.Left != nil || node.Right != nil {
		if node.Left != nil {
			step.Curr = node.Val
			//follow the (*) free to go "down, left"
			step.Next = "DOWN TO THE LEFT"
			*walkThru = append(*walkThru, step)
		} else {
			step.Curr = node.Val
			//follow the (*) free to go "down, left"
			step.Next = "check L, ->"
			*walkThru = append(*walkThru, step)
		}
	} else {
		step.Curr = node.Val
		step.Next = "self-check"
		*walkThru = append(*walkThru, step)
	}

	// fmt.Println("<-, ><")
	postOrderTraversal(node.Left, walkThru, theTraversal)

	// fmt.Println("After go LEFT of: ", node.Val)
	// fmt.Println("->")
	postOrderTraversal(node.Right, walkThru, theTraversal)
	// fmt.Println("After go RIGHT of: ", node.Val)
	// fmt.Println("DONE: ", node.Val)
	*theTraversal = append(*theTraversal, node.Val)
}

func TrackingTreeTraversal(root *modal.TreeNode) {
	var walkThru []modal.WalkThru
	var theTraversal []int

	postOrderTraversal(root, &walkThru, &theTraversal)

	fmt.Println(theTraversal)
	fmt.Println("-----------")
	fmt.Println(walkThru)

}
