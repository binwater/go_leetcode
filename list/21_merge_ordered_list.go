package main

import (
	"fmt"
	"time"
)

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	l3 := &ListNode{}

	if l1.Val > l2.Val{
		l3 = l2
		l3.Next = mergeTwoLists(l1, l2.Next)
	} else {
		l3 = l1
		l3.Next = mergeTwoLists(l1.Next, l2)
	}

	return l3
}

//easy
func main() {
	fmt.Printf("start now %v\n", time.Now().Format("2006-01-02 15:04:05"))
	lastDay := time.Now().AddDate(0, -6, 0)
	fmt.Printf("last is: %v\n", lastDay.Format("2006-01-02 15:04:05"))
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node1.Next = &node2

	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node3.Next = &node4

	//printList(&node3)
	mergeList := mergeTwoLists(&node1, &node3)

	printList(mergeList)
}
