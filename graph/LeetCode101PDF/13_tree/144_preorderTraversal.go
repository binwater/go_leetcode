package main

import "fmt"

/*给你二叉树的根节点 root ，返回它节点值的 前序 遍历。



示例 1：

输入：root = [1,null,2,3]
输出：[1,2,3]

示例 2：

输入：root = []
输出：[]

示例 3：

输入：root = [1]
输出：[1]

示例 4：

输入：root = [1,2]
输出：[1,2]

示例 5：

输入：root = [1,null,2]
输出：[1,2]



提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100



进阶：递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) (res []int) {
	var p func(root *TreeNode)
	p = func(node *TreeNode) {
		if node==nil {
			return
		}
		res = append(res,node.Val)
		p(node.Left)
		p(node.Right)
	}
	p(root)
	return
}

func main() {
	tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{3, nil, nil}

	tree3.Left = &tree1
	tree3.Right = &tree2
	ret := preorderTraversal(&tree3)
	fmt.Println("ret is: ", ret)
}