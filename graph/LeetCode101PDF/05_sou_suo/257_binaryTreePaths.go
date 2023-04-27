package main

import (
	"fmt"
	"strconv"
)

/*给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

叶子节点 是指没有子节点的节点。


示例 1：

输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]

示例 2：

输入：root = [1]
输出：["1"]



提示：

树中节点的数目在范围 [1, 100] 内
-100 <= Node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*var paths []string
func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	treeDfs(root, "")
	return paths
}

func treeDfs(root *TreeNode, ret string) {
	if root != nil{
		tmpPath := ret
		tmpPath += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, tmpPath)
		} else {
			tmpPath += "->"
			treeDfs(root.Left, tmpPath)
			treeDfs(root.Right, tmpPath)
		}
	}
}*/

func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}

	var nodeQueue []*TreeNode
	var pathsQueue []string

	nodeQueue = append(nodeQueue, root)
	pathsQueue = append(pathsQueue, strconv.Itoa(root.Val))

	for i := 0; i < len(nodeQueue); i++ {
		tmpNode, tmpPath := nodeQueue[i], pathsQueue[i]

		if tmpNode.Left == nil && tmpNode.Right == nil{
			paths = append(paths, tmpPath)
			continue
		}


		if tmpNode.Left != nil {
			nodeQueue = append(nodeQueue, tmpNode.Left)
			pathsQueue = append(pathsQueue, tmpPath+"->"+strconv.Itoa(tmpNode.Left.Val))
		}

		if tmpNode.Right != nil {
			nodeQueue = append(nodeQueue, tmpNode.Right)
			pathsQueue = append(pathsQueue, tmpPath+"->"+strconv.Itoa(tmpNode.Right.Val))
		}
	}

	return paths
}

func main() {
	node1 := TreeNode{5, nil, nil}
	node2 := TreeNode{2, nil, &node1}
	node3 := TreeNode{3, nil, nil}
	nodeHead := TreeNode{1, &node2, &node3}

	ret := binaryTreePaths(&nodeHead)
	fmt.Println("ret is: ", ret)
}
