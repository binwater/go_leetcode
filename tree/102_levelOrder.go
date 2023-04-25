package main

import "fmt"

//type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	var ret [][]int
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		ret = append(ret, []int{})
		for count := len(queue); count > 0; count-- {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			ret[level] = append(ret[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}

	return ret
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
