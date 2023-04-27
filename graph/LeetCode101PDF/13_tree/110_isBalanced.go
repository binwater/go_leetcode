package main

import (
	"fmt"
	"math"
)

/*给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。



示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：true

示例 2：

输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

示例 3：

输入：root = []
输出：true



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/balanced-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil{
		return 0
	}

	left := helper(root.Left)
	right := helper(root.Right)
	if left == -1 || right == -1 || math.Abs(float64(left-right))>1{
		return -1
	}

	return 1+int(math.Max(float64(left), float64(right)))
}

func main2()  {
	tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{3, nil, nil}
	tree4 := TreeNode{4, nil, nil}
	tree4.Left = &tree3
	tree4.Right = &tree2
	tree3.Left = &tree1
	ret := isBalanced(&tree4)
	fmt.Println("ret is: ", ret)
}
