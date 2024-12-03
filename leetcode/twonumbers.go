package leetcode

//Add Two Numbers I

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]

func _() {
	l1, l2 := []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}
	linkedl17 := &ListNode1{Val: l1[6], Next: nil}
	linkedl16 := &ListNode1{Val: l1[5], Next: linkedl17}
	linkedl15 := &ListNode1{Val: l1[4], Next: linkedl16}
	linkedl14 := &ListNode1{Val: l1[3], Next: linkedl15}
	linkedl13 := &ListNode1{Val: l1[2], Next: linkedl14}
	linkedl12 := &ListNode1{Val: l1[1], Next: linkedl13}
	head1 := &ListNode1{Val: l1[0], Next: linkedl12}

	linkedl24 := &ListNode1{Val: l2[3], Next: nil}
	linkedl23 := &ListNode1{Val: l2[2], Next: linkedl24}
	linkedl22 := &ListNode1{Val: l2[1], Next: linkedl23}
	head2 := &ListNode1{Val: l2[0], Next: linkedl22}

	addTwoNumbers(head1, head2)
}

// Definition for singly-linked list.
type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func addTwoNumbers(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	list1 := []*ListNode1{l1}
	listvalues1 := []int{l1.Val}

	list2 := []*ListNode1{l2}
	listvalues2 := []int{l2.Val}

	current := list1[0]
	for current.Next != nil {
		list1 = append(list1, current.Next)
		listvalues1 = append(listvalues1, current.Next.Val)
		current = current.Next
	}

	current = list2[0]
	for current.Next != nil {
		list2 = append(list2, current.Next)
		listvalues2 = append(listvalues2, current.Next.Val)
		current = current.Next
	}

	biglist := list1
	minlist := list2
	if len(list2) > len(list1) {
		biglist = list2
		minlist = list1
	}

	lastIndex := 0
	upper := 0
	for i := 0; i < len(minlist); i++ {
		new := biglist[i].Val + minlist[i].Val + upper
		upper = 0
		if new >= 10 {
			upper++
			new = new % 10
		}
		biglist[i].Val = new
		lastIndex = i

	}
	for upper > 0 {
		if lastIndex < len(biglist)-1 {
			new := biglist[lastIndex+1].Val + upper
			upper = 0
			if new >= 10 {
				new = new % 10
				upper++
			}
			biglist[lastIndex+1].Val = new
			lastIndex++
		} else {
			newelement := &ListNode1{Val: 1, Next: nil}
			biglist[len(biglist)-1].Next = newelement
			upper--
		}
	}
	return biglist[0]
}
