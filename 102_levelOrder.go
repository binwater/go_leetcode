package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res,[]int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}

func main() {
	node1 := &TreeNode{3, nil, nil}
	node1.Left = &TreeNode{9, nil, nil}
	node1.Right = &TreeNode{20, nil, nil}
	node1.Right.Left = &TreeNode{15, nil, nil}
	node1.Right.Right = &TreeNode{7, nil, nil}

	res := levelOrder(node1)
	fmt.Printf("res is: %v\n", res)
}