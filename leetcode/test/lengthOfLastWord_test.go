package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	words := []string{"Hello World", "   fly me   to   the moon  ",
		"luffy is still joyboy", "a "}
	expected := []int{5, 4, 6, 1}
	for x := 0; x < len(words); x++ {
		fmt.Println("=======================================")
		output := leetcode.LengthOfLastWord(words[x])
		if output != expected[x] {
			fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
			continue
		}
		fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
