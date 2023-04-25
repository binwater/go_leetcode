package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(maxDepth(root.Left), maxDepth(root.Left)) + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
