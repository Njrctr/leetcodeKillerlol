package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {

	first, second := []string{"11", "1010", "1", "101111"}, []string{"1", "1011", "111", "10"}
	expected := []string{"100", "10101", "1000", "110001"}
	for x := 0; x < len(first); x++ {
		fmt.Println("=======================================")
		output := leetcode.AddBinary(first[x], second[x])
		if output != expected[x] {
			fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
			continue
		}
		fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
