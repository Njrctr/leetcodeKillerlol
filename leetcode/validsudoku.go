package leetcode

import (
	"bytes"
)

func IsValidSudoku(board [][]byte) bool {
	for _, line := range board {
		for index, i := range line {
			if i == 46 {
				continue
			}
			if bytes.Count(line, []byte{i}) > 1 {
				return false
			}
		}

	}
	for x := 0; x < 9; x++ {
		// for col := range
	}
	return false
}
