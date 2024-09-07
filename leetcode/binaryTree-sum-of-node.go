package leetcode

import (
	"fmt"
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

func nodeRecursive(root *modal.TreeNode, overallTreePath *int) int {
	if root == nil {
		return 0
	}

	var perfectThree Result
	var left *int
	var right *int
	perfectThree.Path = append(perfectThree.Path, root.Val)
	perfectThree.Sum = root.Val

	//FROM: left
	maxLeftPath := max(nodeRecursive(root.Left, overallTreePath), 0)
	if root.Left != nil {
		left = &root.Left.Val
		perfectThree.Path = append(perfectThree.Path, *left)
		perfectThree.Sum += *left

	}

	//To: Right
	maxRightPath := max(nodeRecursive(root.Right, overallTreePath), 0)
	if root.Right != nil {
		right = &root.Right.Val
		perfectThree.Path = append(perfectThree.Path, *right)
		perfectThree.Sum += *right
	}

	// To: root
	currentNodePathMax := root.Val + maxLeftPath + maxRightPath
	*overallTreePath = max(*overallTreePath, currentNodePathMax)

	return root.Val + max(maxLeftPath, maxRightPath)
}

// postorder traversal left->right->root
func MaxPathSum(root *modal.TreeNode) int {
	overallTreePath := -1 << 63
	nodeRecursive(root, &overallTreePath)
	fmt.Println(overallTreePath)
	return overallTreePath
}
