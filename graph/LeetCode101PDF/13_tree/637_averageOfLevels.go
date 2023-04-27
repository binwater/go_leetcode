package main

import "fmt"

/*给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。



示例 1：

输入：root = [3,9,20,null,null,15,7]
输出：[3.00000,14.50000,11.00000]
解释：第 0 层的平均值为 3,第 1 层的平均值为 14.5,第 2 层的平均值为 11 。
因此返回 [3, 14.5, 11] 。

示例 2:

输入：root = [3,9,20,15,7]
输出：[3.00000,14.50000,11.00000]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/average-of-levels-in-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	ret := make([]float64, 0)

	var queue = []*TreeNode{root}
	var level int

	for len(queue) > 0 {
		counter := len(queue)
		tmpArr := make([]int, 0)
		for counter > 0{
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			tmpArr = append(tmpArr, queue[0].Val)
			queue = queue[1:]
		}
		level++
		tmpAll := 0
		for _, v := range tmpArr{
			tmpAll += v
		}
		ret = append(ret, float64(tmpAll)/float64(len(tmpArr)))
	}
	return ret
}

func main() {
	tree1 := TreeNode{1, nil, nil}
	tree2 := TreeNode{2, nil, nil}
	tree3 := TreeNode{2, nil, nil}
	tree4 := TreeNode{4, nil, nil}
	tree4.Left = &tree3
	tree4.Right = &tree2
	tree3.Left = &tree1
	ret := averageOfLevels(&tree4)
	fmt.Println("ret is: ", ret)
}
