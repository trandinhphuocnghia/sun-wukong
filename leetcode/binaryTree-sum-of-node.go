package leetcode

import (
	"cmp"
	"fmt"
	"slices"
	"sunwukong/modal"
)

/*
THE 1st approach:
- Algorithm: The postorder traversal.
- Dimension of the recursive loop: Left->Right->Root.
*/
type Result struct {
	Sum  int
	Path []int
}

// recursive func to loop through all node
func postOrderTraversal(root *modal.TreeNode, result *[]Result) {
	if root == nil {
		return
	}

	var perfectThree Result
	var left *int
	var right *int
	perfectThree.Path = []int{}
	perfectThree.Sum = 0

	//FROM: left
	postOrderTraversal(root.Left, result)
	if root.Left != nil {
		left = &root.Left.Val
		perfectThree.Path = append(perfectThree.Path, *left)
		fmt.Println("push LEFT: ", root.Left.Val)
		if *left >= 0 {
			perfectThree.Sum += *left
		}
	}

	//To: Right
	postOrderTraversal(root.Right, result)
	if root.Right != nil {
		right = &root.Right.Val
		perfectThree.Path = append(perfectThree.Path, *right)
		fmt.Println("push RIGHT: ", root.Right.Val)
		if *right >= 0 {
			perfectThree.Sum += *right
		}
	}
	// To: root
	perfectThree.Path = append(perfectThree.Path, root.Val)
	perfectThree.Sum += root.Val
	*result = append(*result, perfectThree)
	fmt.Println("node: ", root.Val)
}

// inOrder traversal to find the streak
func inOrderTraversal(root *modal.TreeNode, result *int) {
	if root == nil {
		return
	}

	// var left *int
	// var right *int

	//from left
	inOrderTraversal(root.Left, result)
	// if root.Left != nil {
	// 	left = &root.Left.Val
	// 	if *left >= 0 {
	// 		*result += *left
	// 	}
	// }

	//to root
	if root != nil && root.Val >= 0 {
		*result += root.Val
	}

	//to right
	inOrderTraversal(root.Right, result)
	if root.Right == nil {
		if root != nil && root.Val >= 0 {
			*result -= root.Val
		}
	}
}

func maxPath(root *modal.TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	leftPathSum := maxPath(root.Left, ans)
	rightPathSum := maxPath(root.Right, ans)
	*ans = max(*ans, leftPathSum+rightPathSum+root.Val)
	return max(max(leftPathSum+root.Val, rightPathSum+root.Val), 0)
}

func MaxPathSum(root *modal.TreeNode) int {
	var result []Result

	// postorder traversal left->right->root
	postOrderTraversal(root, &result)
	maxOfTree := slices.MaxFunc(result, func(i, j Result) int {
		return cmp.Compare(i.Sum, j.Sum)
	})

	return maxOfTree.Sum
}
