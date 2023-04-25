package main

import "math"

var pre = math.MinInt32

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
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST1(root.Left) {
		return false
	}

	if root.Val <= pre {
		return false
	}

	pre = root.Val
	return isValidBST1(root.Right)
}

func isValidBST(root *TreeNode) bool {

	var dfs func(root *TreeNode, min, max int) bool
	dfs = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}

		if root.Val <= min || root.Val >= max {
			return false
		}

		return dfs(root.Left, min, root.Val) && dfs(root.Right, root.Val, max)
	}

	return dfs(root, math.MinInt32, math.MaxInt32)
}

func main() {

}
