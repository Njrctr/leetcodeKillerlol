package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestIsPrefixOfWord(t *testing.T) {

	first := []string{"i am tired", "i love eating burger", "this problem is an easy problem", "hello from the other side"}
	second := []string{"you", "burg", "pro", "they"}
	expected := []int{-1, 4, 2, -1}
	for x := 0; x < len(first); x++ {
		fmt.Println("=======================================")
		output := leetcode.IsPrefixOfWord(first[x], second[x])
		if output != expected[x] {
			fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
			continue
		}
		fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
