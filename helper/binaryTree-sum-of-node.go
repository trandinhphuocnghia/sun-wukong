package helper

import (
	"fmt"
	"sunwukong/modal"
)

/*
* Input: Array
* Output: TreeNode
 */
func ArrayToTree(arr []interface{}) *modal.TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	//root of tree is the first element
	root := &modal.TreeNode{Val: arr[0].(int)}

	//for each node will push three elements for: node, left, right
	subTree := []*modal.TreeNode{root}
	i := 1

	for i < len(arr) {
		current := subTree[0]
		subTree = subTree[1:]

		fmt.Println("node", current.Val)
		fmt.Println("subTree", subTree)
		if i < len(arr) && arr[i] != nil {
			current.Left = &modal.TreeNode{Val: arr[i].(int)}
			fmt.Println("LEFT:", arr[i].(int))
			subTree = append(subTree, current.Left)
			fmt.Println("Left, subTree", subTree)
		}
		i++

		if i < len(arr) && arr[i] != nil {
			current.Right = &modal.TreeNode{Val: arr[i].(int)}
			fmt.Println("RIGHT:", arr[i].(int))
			subTree = append(subTree, current.Right)
			fmt.Println("Right, subTree", subTree)
		}
		i++
	}

	return root
}
