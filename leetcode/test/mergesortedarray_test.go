package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := [][]int{{1, 2, 3, 0, 0, 0}, {1}}
	nums2 := [][]int{{2, 5, 6}, {}}
	m, n := []int{3, 1}, []int{3, 0}
	expected := [][]int{{1, 2, 2, 3, 5, 6}, {1}}
	for x := 0; x < len(nums1); x++ {
		fmt.Println("=======================================")
		output := leetcode.Merge(nums1[x], m[x], nums2[x], n[x])
		fmt.Println("Ожидалось:", expected[x], "Вывод теста:", output)
		// if output != expected[x] {
		// 	fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
		// 	continue
		// }
		// fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
