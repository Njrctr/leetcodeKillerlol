package leetcode

import (
	"alg/tracking"
	"fmt"
)

func _() {
	n := 5
	fmt.Println(countAndSay(n))
	fmt.Println(countAndSayByte(n))
}

func countAndSay(n int) string {
	defer tracking.Duration(tracking.Track("countAndSay: "))
	if n == 1 {
		return "1"
	}
	start := []byte{'1'}
	for x := 2; x <= n; x++ {

		lastelem := start[0]
		count := []byte{'0'}
		newarr := []byte{}
		for el := 0; el < len(start); el++ { // Подсчёт элементов
			if start[el] != lastelem {
				newarr = append(newarr, []byte{count[0], lastelem}...)
				lastelem = start[el]
				count = []byte{'1'}
			} else {
				count[0]++
			}
		}
		newarr = append(newarr, []byte{count[0], lastelem}...)
		start = newarr
	}
	return string(start)
}

func countAndSayByte(n int) string {
	defer tracking.Duration(tracking.Track("countAndSayByte: "))
	start := []byte{'1'}
	for n > 1 {
		newstart := []byte{'0', start[0]}
		for _, cur := range start {
			if cur == newstart[len(newstart)-1] {
				newstart[len(newstart)-2]++
				continue
			}
			newstart = append(newstart, []byte{'1', cur}...)
		}
		start = newstart
		n--
	}
	return string(start)
}
