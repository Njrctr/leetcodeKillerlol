package search

import (
	"fmt"
)

func Binary(search int, list []int) {
	low := 0
	top := len(list) - 1
	for x := 0; low <= top; x++ {
		now := (low + top) / 2
		if list[now] > search {
			top = now - 1
			continue
		}
		if list[now] < search {
			low = now + 1
			continue
		}
		fmt.Printf("[+++]Найдено: %d\nКоличество итераций: %d\n", now, x)
		return
	}

	fmt.Println("[---]Элемент не обнаружен!")
}
