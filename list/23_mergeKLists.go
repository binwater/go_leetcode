package main

//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//示例 2：
//
//输入：lists = []
//输出：[]
//
//示例 3：
//
//输入：lists = [[]]
//输出：[]
//
//
//
//提示：
//
//k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	result := lists[0]
	for i:=1; i<len(lists); i++{
		result = mergeTwo(result, lists[i])
	}

	return result
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var l3 *ListNode

	if l1.Val > l2.Val{
		l3 = l2
		l3.Next = mergeTwo(l1, l2.Next)
	} else {
		l3 = l1
		l3.Next = mergeTwo(l1.Next, l2)
	}

	return l3
}

//hard
func main() {
	node := ListNode{1, nil}
	node.Next = &ListNode{4, nil}
	node.Next.Next = &ListNode{6, nil}

	node1 := ListNode{3, nil}
	node1.Next = &ListNode{7, nil}
	node1.Next.Next = &ListNode{8, nil}

	node2 := ListNode{2, nil}
	node2.Next = &ListNode{5, nil}

	kLists := make([]*ListNode, 0)
	kLists = append(kLists, &node, &node1, &node2)

	after := mergeKLists(kLists)

	printList(after)
}