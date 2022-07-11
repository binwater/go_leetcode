package main

import "math"

var pre = math.MinInt32
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left){
		return false
	}

	if root.Val <= pre{
		return false
	}

	pre = root.Val
	return isValidBST(root.Right)
}


func main() {

}