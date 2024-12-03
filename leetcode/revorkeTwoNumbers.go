package leetcode

//Add Two Numbers II

func _() {
	l1, l2 := []int{3, 3, 4, 3, 2}, []int{5, 6, 4}
	linkedl15 := &ListNode2{Val: l1[4], Next: nil}
	linkedl14 := &ListNode2{Val: l1[3], Next: linkedl15}
	linkedl13 := &ListNode2{Val: l1[2], Next: linkedl14}
	linkedl12 := &ListNode2{Val: l1[1], Next: linkedl13}
	linkedl11 := &ListNode2{Val: l1[0], Next: linkedl12}

	linkedl23 := &ListNode2{Val: l2[2], Next: nil}
	linkedl22 := &ListNode2{Val: l2[1], Next: linkedl23}
	linkedl21 := &ListNode2{Val: l2[0], Next: linkedl22}

	// addTwoNumbers(linkedl11, linkedl21)
	revorkeTwoNumbers2(linkedl11, linkedl21)
}

// Definition for singly-linked list.
type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func revorkeTwoNumbers2(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	list1 := []*ListNode2{l1}
	list1total := []int{l1.Val}

	list2 := []*ListNode2{l2}
	list2total := []int{l2.Val}

	current := l1
	for current.Next != nil {
		list1 = append(list1, current.Next)
		list1total = append(list1total, current.Next.Val)
		current = current.Next

	}
	current = l2
	for current.Next != nil {
		list2 = append(list2, current.Next)
		list2total = append(list2total, current.Next.Val)
		current = current.Next
	}
	bigl := list2
	biglist := list2total
	smalllist := list1total
	if len(list2) < len(list1) {
		bigl = list1
		biglist = list1total
		smalllist = list2total
	}
	neededlist := twoListsFuck(biglist, smalllist)
	needup := len(neededlist) > len(bigl)
	for i := 1; i <= len(bigl); i++ {
		bigl[len(bigl)-i].Val = neededlist[len(neededlist)-i]
	}
	if needup {
		mewelement := &ListNode2{Val: 1, Next: bigl[0]}
		return mewelement
	}
	return bigl[0]
}

func twoListsFuck(big, small []int) []int {
	needup := 0
	lastIndex := 0
	for x := 1; x <= len(small); x++ {
		new := big[len(big)-x] + small[len(small)-x] + needup
		if new >= 10 {
			new = new % 10
			needup = 1
		} else {
			needup = 0
		}
		big[len(big)-x] = new
		lastIndex = len(big) - x
	}
	if needup != 0 {
		for needup > 0 {
			if lastIndex != 0 {
				new := big[lastIndex-1] + needup
				if new < 10 {
					needup--
				} else {
					new = new % 10
				}
				big[lastIndex-1] = new
				lastIndex--
			} else {
				newarr := []int{1}
				newarr = append(newarr, big...)
				big = newarr
				needup--
			}
		}
	}
	return big
}
