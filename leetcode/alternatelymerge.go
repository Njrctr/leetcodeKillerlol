package leetcode

import "fmt"

func _() {
	word1, word2 := "ab", "pqrs"
	fmt.Println(mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	bytes1, bytes2 := []byte(word1), []byte(word2)
	l1, l2 := len(bytes1), len(bytes2)
	var biggest *[]byte
	minlen := l1
	biggest = &bytes2
	if l1 >= l2 {
		biggest = &bytes1
		minlen = l2
	}
	result := []byte{}
	i := 0
	for i < minlen {
		result = append(result, []byte{bytes1[i], bytes2[i]}...)
		i++
	}
	if l1 != l2 {
		result = append(result, (*biggest)[i:]...)
	}
	return string(result)
}
