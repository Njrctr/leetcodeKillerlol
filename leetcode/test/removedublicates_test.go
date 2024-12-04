package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	l3 := &leetcode.ListNode{Val: 2, Next: nil}
	l2 := &leetcode.ListNode{Val: 1, Next: l3}
	l1 := &leetcode.ListNode{Val: 1, Next: l2}

	r5 := &leetcode.ListNode{Val: 3, Next: nil}
	r4 := &leetcode.ListNode{Val: 3, Next: r5}
	r3 := &leetcode.ListNode{Val: 2, Next: r4}
	r2 := &leetcode.ListNode{Val: 1, Next: r3}
	r1 := &leetcode.ListNode{Val: 1, Next: r2}

	t3 := &leetcode.ListNode{Val: 1, Next: nil}
	t2 := &leetcode.ListNode{Val: 1, Next: t3}
	t1 := &leetcode.ListNode{Val: 1, Next: t2}

	heads := []*leetcode.ListNode{l1, r1, t1}
	expected := [][]int{{1, 2}, {1, 2, 3}, {1}}
	for x := 0; x < len(heads); x++ {
		fmt.Println("=======================================")
		result := leetcode.DeleteDuplicates(heads[x])
		var resultList []int
		for node := result; node != nil; node = node.Next {
			resultList = append(resultList, node.Val)
		}
		fmt.Println("Ожидалось:", expected[x], "Вывод теста:", resultList)
		// if output != expected[x] {
		// 	fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
		// 	continue
		// }
		// fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
