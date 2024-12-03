package leetcode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		memoNext := head.Next
		head.Next = prev
		prev = head
		head = memoNext
	}
	return prev
}
