package leetcode

// Definition for singly-linked list.

func DeleteDuplicates(head *ListNode) *ListNode {

	node := head
	for node != nil {
		for node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}
