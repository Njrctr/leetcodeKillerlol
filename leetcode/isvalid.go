package leetcode

import "strings"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	symbols := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",

		"}": "{",
		"]": "[",
		")": "(",
	}
	var vanga []string
	open := "{[("
	for _, sym := range s {
		if strings.ContainsRune(open, sym) {
			vanga = append(vanga, string(sym))
		} else {
			if len(vanga) < 1 {
				return false
			}
			if symbols[string(sym)] == vanga[len(vanga)-1] {
				vanga = vanga[:len(vanga)-1]
			} else {
				return false
			}
		}
	}
	if len(vanga) != 0 {
		return false
	}
	return true
}
