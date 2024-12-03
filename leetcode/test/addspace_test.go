package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestAddSpace(t *testing.T) {

	first := []string{"LeetcodeHelpsMeLearn", "icodeinpython", "spacing"}
	second := [][]int{{8, 13, 15}, {1, 5, 7, 9}, {0, 1, 2, 3, 4, 5, 6}}
	expected := []string{"Leetcode Helps Me Learn", "i code in py thon", " s p a c i n g"}
	for x := 0; x < len(first); x++ {
		fmt.Println("=======================================")
		output := leetcode.AddSpace(first[x], second[x])
		if output != expected[x] {
			fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
			continue
		}
		fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
