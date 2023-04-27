package main

import (
	"fmt"
	"math"
)

/*给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。



示例 :
给定二叉树

1
/ \
2   3
/ \
4   5

返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。



注意：两结点之间的路径长度是以它们之间边的数目表示。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var diameter = 0
	helperDiameter(root, &diameter)
	return diameter
}

func helperDiameter(node *TreeNode, diameter *int) int {
	if node == nil{
		return 0
	}

	l := helperDiameter(node.Left, diameter)
	r := helperDiameter(node.Right, diameter)

	*diameter = int(math.Max(float64(l+r), float64(*diameter)))
	return int(math.Max(float64(l), float64(r)))+1
}

func main3() {
	tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{3, nil, nil}
	tree4 := TreeNode{4, nil, nil}
	tree4.Left = &tree3
	tree4.Right = &tree2
	tree3.Left = &tree1
	ret := diameterOfBinaryTree(&tree4)
	fmt.Println("ret is: ", ret)
}
