package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {

	first := []string{"PAYPALISHIRING", "PAYPALISHIRING", "A"}
	second := []int{3, 4, 1}
	expected := []string{"PAHNAPLSIIGYIR", "PINALSIGYAHRPI", "A"}
	for x := 0; x < len(first); x++ {
		fmt.Println("=======================================")
		output := leetcode.Convert(first[x], second[x])
		if output != expected[x] {
			fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
			continue
		}
		fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
