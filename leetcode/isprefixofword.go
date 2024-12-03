package leetcode

import (
	"fmt"
	"strings"
)

func IsPrefixOfWord(sentence string, searchWord string) int {
	sliceStr := strings.Split(sentence, " ")
	byteSearchWord := []byte(searchWord)
	for i, word := range sliceStr {
		byteWord := []byte(word)
		if len(byteWord) < len(byteSearchWord) {
			continue
		}
		var skip bool = false
		for x := 0; x < len(byteSearchWord); x++ {
			fmt.Println(byteWord[x], byteSearchWord[x])
			if byteWord[x] != byteSearchWord[x] {
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		return i + 1
	}
	return -1
}
