package main

import (
	"fmt"
)

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？
//
//
//
//示例 1：
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headTmp := ListNode{0,head}

	firstNode := &headTmp
	secondNode := &headTmp

	//两个node之前差值为n
	for i:=1; i<=n+1; i++{
		firstNode = firstNode.Next
	}

	for firstNode != nil {
		firstNode = firstNode.Next
		secondNode = secondNode.Next
	}

	secondNode.Next = secondNode.Next.Next

	return headTmp.Next
}

func printList(list *ListNode) {
	for list.Next != nil {
		fmt.Printf("value is: %v\n", list.Val)
		list = list.Next
	}
	fmt.Printf("value is: %v\n", list.Val)
}

func main1() {
	node := ListNode{1, nil}
	node.Next = &ListNode{2, nil}
	node.Next.Next = &ListNode{3, nil}
	node.Next.Next.Next = &ListNode{4, nil}

	printList(&node)
	after := removeNthFromEnd(&node, 3)
	fmt.Println("after delete")
	printList(after)
}