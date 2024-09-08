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

func nodeRecursive(root *modal.TreeNode, overallTreePath *int, path *[]int) (int, []int) {
	if root == nil {
		return 0, []int{}
	}

	//FROM: left
	maxLeftPath, leftPath := nodeRecursive(root.Left, overallTreePath, path)
	if maxLeftPath < 0 {
		maxLeftPath = 0
		leftPath = []int{}
	}

	//To: Right
	maxRightPath, rightPath := nodeRecursive(root.Right, overallTreePath, path)
	if maxRightPath < 0 {
		maxRightPath = 0
		rightPath = []int{}
	}

	currentNodePathMax := root.Val + maxLeftPath + maxRightPath

	//result, check current subtree with overall if it's gt so that is the result.
	if currentNodePathMax > *overallTreePath {
		*overallTreePath = currentNodePathMax
		*path = append(leftPath, root.Val)
		*path = append(*path, rightPath...)
	}
	fmt.Println("---")
	fmt.Println("node", root.Val)
	fmt.Println("leftPath", leftPath)
	fmt.Println("rightPath", rightPath)
	fmt.Println("++++++++++++++++++++++")

	if maxLeftPath > maxRightPath {
		return root.Val + maxLeftPath, append(leftPath, root.Val)
	} else {
		return root.Val + maxRightPath, append(rightPath, root.Val)
	}

}

// postorder traversal left->right->root
func MaxPathSum(root *modal.TreeNode) int {
	overallTreePath := -1 << 63
	var path []int
	//var Paths []int
	nodeRecursive(root, &overallTreePath, &path)
	fmt.Println(overallTreePath)
	fmt.Println(path)
	return overallTreePath
}
