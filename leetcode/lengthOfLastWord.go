package leetcode

import (
	"fmt"
	"strings"
)

func LengthOfLastWord(s string) int {
	splitedStr := strings.Split(s, "")
	fmt.Println(len(splitedStr))
	for {
		if splitedStr[len(splitedStr)-1] == " " {
			splitedStr = splitedStr[:len(splitedStr)-1]
			continue
		}
		break
	}
	str := strings.Join(splitedStr, "")
	return len(str) - 1 - strings.LastIndex(str, " ")
}
