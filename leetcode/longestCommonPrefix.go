package leetcode

import "bytes"

func longestCommonPrefix(strs []string) string {
	minimalStr := strs[0]
	for x := 1; x < len(strs); x++ {
		if len(minimalStr) > len(strs[x]) {
			minimalStr = strs[x]
		}
	}
	var prefix bytes.Buffer
	for i := 0; i < len(minimalStr); i++ {
		to_add := true
		for j := 0; j < len(strs); j++ {
			if strs[j] == minimalStr {
				continue
			}
			if strs[j][i] != minimalStr[i] {
				to_add = false
			}
		}
		if to_add {
			prefix.WriteByte(minimalStr[i])
		} else {
			break
		}
	}
	return prefix.String()
}
