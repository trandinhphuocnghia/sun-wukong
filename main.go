package main

import (
	"fmt"
	"sunwukong/leetcode"
	"sunwukong/modal"
)

func main() {
	fmt.Println("THE GREAT SAGE")

	//1. input is array ?
	//covert
	//fmt.Println(helper.ArrayToTree([]interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}))

	//2. input is in TreeNode struct
	v := modal.TreeNode{
		Val: 5,
		Left: &modal.TreeNode{
			Val: 4,
			Right: &modal.TreeNode{
				Val: 11,
				Left: &modal.TreeNode{
					Val: 7,
				},
				Right: &modal.TreeNode{
					Val: 2,
				},
			},
		},
		Right: &modal.TreeNode{
			Val: 8,
			Left: &modal.TreeNode{
				Val: 13,
			},
			Right: &modal.TreeNode{
				Val: 4,
				Left: &modal.TreeNode{
					Val: 1,
				},
			},
		},
	}
	leetcode.MaxPathSum(&v)
}
