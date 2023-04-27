package main

import "fmt"

/*给你一个二叉树的根节点 root ， 检查它是否轴对称。



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true

示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/symmetric-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSymmetricHelp(root.Left, root.Right)
}

func isSymmetricHelp(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricHelp(left.Left, right.Right) &&
		isSymmetricHelp(left.Right, right.Left)
}

func main5() {
	//tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{2, nil, nil}
	tree4 := TreeNode{4, nil, nil}
	tree4.Left = &tree3
	tree4.Right = &tree2
	//tree3.Left = &tree1
	ret := isSymmetric(&tree4)
	fmt.Println("ret is: ", ret)
}
