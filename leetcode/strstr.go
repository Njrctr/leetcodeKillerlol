package leetcode

import "fmt"

func strStr(haystack string, needle string) int {
	needleRunes := []rune(needle)
	haystackRunes := []rune(haystack)

	for x := 0; x < len(haystackRunes); x++ {
		if haystackRunes[x] == needleRunes[0] {
			fmt.Println("x: ", x, "\nlen(needleRunes): ", len(needleRunes), "\nlen(haystackRunes): ", len(haystackRunes))
			if x+len(needleRunes) > len(haystackRunes) {
				return -1
			}
			slice := haystackRunes[x : x+len(needleRunes)]
			fmt.Println("slice: ", slice)
			if string(slice) == string(needleRunes) {
				return x
			}
		}
	}
	return -1
}
