package main

import "fmt"

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode2) *TreeNode2 {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		head.Left, head.Right = head.Right, head.Left
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
	}
	return root
}

func main() {
	root := TreeNode2{1, nil, nil}
	root.Left = &TreeNode2{2, nil, nil}
	root.Right = &TreeNode2{3, nil, nil}
	fmt.Println(invertTree(&root))

}
