package search

import (
	"fmt"
)

func Lazy(search int, list []int) {
	for x := 0; x <= len(list)-1; x++ {
		now := list[x]
		if now != search {
			continue
		}
		fmt.Printf("[+++]Найдено: %d\nКоличество итераций: %d\n", now, x)
		return
	}
	fmt.Println("[---]Элемент не обнаружен!")
}
