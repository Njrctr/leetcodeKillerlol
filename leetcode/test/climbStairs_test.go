package leetcode_test

import (
	"alg/leetcode"
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	steps := []int{2, 3, 4}
	expected := []int{2, 3, 5}
	for x := 0; x < len(steps); x++ {
		fmt.Println("=======================================")
		output := leetcode.ClimbStairs(steps[x])
		if output != expected[x] {
			fmt.Println("[---]Ошибка кейса №", x+1, "Ожидалось:", expected[x], " Получено:", output)
			continue
		}
		fmt.Println("[+++]Кейс №", x+1, "Выполнен успешно!")
		fmt.Println("=======================================")
	}
}
