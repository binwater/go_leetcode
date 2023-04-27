package main

import "fmt"

/*给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。



示例 1：

输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。

示例 2：

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/path-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	hash := map[int]int{0:1}

	var dfs func(root *TreeNode, curSum int)
	dfs = func(root *TreeNode, curSum int) {

		if root == nil {
			return
		}
		curSum += root.Val
		if _,ok := hash[curSum-targetSum]; ok {
			res += hash[curSum-targetSum]
		}
		hash[curSum]++

		dfs(root.Left,curSum)
		dfs(root.Right,curSum)
		hash[curSum]--
	}

	dfs(root,0)

	return res

}

func main4() {
	tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{3, nil, nil}
	tree4 := TreeNode{4, nil, nil}
	tree4.Left = &tree3
	tree4.Right = &tree2
	tree3.Left = &tree1
	ret := pathSum(&tree4, 4)
	fmt.Println("ret is: ", ret)
}
