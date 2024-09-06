package main

import (
	"fmt"
	"sunwukong/helper"
	"sunwukong/leetcode"
	"sunwukong/modal"
)

func main() {
	fmt.Println("THE GREAT SAGE")

	//1. out is array ?
	fmt.Println(helper.ArrayToTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}))

	//2. out is TreeNode struct
	v := modal.TreeNode{
		Val: 1,
		Right: &modal.TreeNode{
			Val: 20,
			Left: &modal.TreeNode{
				Val: 15,
			},
			Right: &modal.TreeNode{
				Val: 7,
			},
		},
		Left: &modal.TreeNode{
			Val: -1,
		},
	}
	leetcode.MaxPathSum(&v)
}
