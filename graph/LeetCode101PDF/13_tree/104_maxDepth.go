package main

import (
	"fmt"
	"math"
)

/*给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

3
/ \
9  20
/  \
15   7

返回它的最大深度 3 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//Definition for a binary tree node.
/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/
func maxDepth(root *TreeNode) int {
	if root != nil {
		return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
	}

	return 0
}

func main1()  {
	tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{3, nil, nil}
	tree4 := TreeNode{4, nil, nil}
	tree4.Left = &tree3
	tree4.Right = &tree2
	tree3.Left = &tree1
	ret := maxDepth(&tree4)
	fmt.Println("ret is: ", ret)
}
