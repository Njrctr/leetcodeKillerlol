package leetcode

import (
	"strings"
)

func AddSpace(s string, spaces []int) string {
	slpitedStr := []byte(s)
	var resultList []string
	var left int = 0
	var right int

	for _, space := range spaces {
		right = space
		resultList = append(resultList, string(slpitedStr[left:space]))
		left = right
	}
	resultList = append(resultList, string(slpitedStr[left:]))
	// fmt.Println(strings.Join(resultList, " "))

	return strings.Join(resultList, " ")
}
