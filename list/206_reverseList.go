package main

import "fmt"

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//
//示例 1：
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//
//示例 2：
//
//输入：head = [1,2]
//输出：[2,1]
//
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000
//
//
//
//进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/reverse-linked-list
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	pre := head
	cur := head.Next

	pre.Next = nil
	for cur != nil && cur.Next != nil{
		next := cur.Next

		cur.Next = pre

		pre  = cur
		cur = next
	}

	cur.Next = pre

	return cur
}

//easy
func main() {
	node := ListNode{1, nil}
	//node.Next = &ListNode{2, nil}
	//node.Next.Next = &ListNode{3, nil}
	//node.Next.Next.Next = &ListNode{4, nil}
	printList(&node)
	fmt.Println("after reverse ")
	after := reverseList(&node)
	printList(after)
}