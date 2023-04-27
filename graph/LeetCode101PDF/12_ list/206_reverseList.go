package main

import "fmt"

/*给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。



示例 1：

输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：

输入：head = [1,2]
输出：[2,1]

示例 3：

输入：head = []
输出：[]



来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}

	pre, cur := head, head.Next
	pre.Next = nil
	for cur.Next != nil{
		next := cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	cur.Next = pre
	return cur
}

func reverseList(head *ListNode) *ListNode {
	if head==nil || head.Next==nil{
		return head
	}

	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next=nil

	return ret
}

func ListPrint(head *ListNode){
	for head.Next!=nil{
		fmt.Println("node val is: ", head.Val)
		head = head.Next
	}
	fmt.Println("node val is: ", head.Val)
}

func main() {
	node1 := ListNode{1, nil}
	node2 := ListNode{2, nil}
	node3 := ListNode{3, nil}
	node4 := ListNode{4, nil}
	node4.Next = &node3
	node3.Next = &node2
	node2.Next = &node1
	ListPrint(&node4)

	fmt.Println("")
	after := reverseList(&node4)
	ListPrint(after)
}
