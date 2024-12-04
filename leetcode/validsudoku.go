package leetcode

import (
	"bytes"
)

func IsValidSudoku(board [][]byte) bool {
	var boxes [9][]byte
	for indexLine, line := range board { // Проверка всех линий на повторяющиеся символы
		for indexCol, i := range line {
			if i == '.' {
				continue
			}

			if indexLine < 3 { // Сборка боксов 3х3
				if indexCol < 3 {
					boxes[0] = append(boxes[0], i)
				} else if indexCol >= 3 && indexCol < 6 {
					boxes[1] = append(boxes[1], i)
				} else if indexCol >= 6 {
					boxes[2] = append(boxes[2], i)
				}
			} else if indexLine >= 3 && indexLine < 6 {
				if indexCol < 3 {
					boxes[3] = append(boxes[3], i)
				} else if indexCol >= 3 && indexCol < 6 {
					boxes[4] = append(boxes[4], i)
				} else if indexCol >= 6 {
					boxes[5] = append(boxes[5], i)
				}
			} else if indexLine >= 6 {
				if indexCol < 3 {
					boxes[6] = append(boxes[6], i)
				} else if indexCol >= 3 && indexCol < 6 {
					boxes[7] = append(boxes[7], i)
				} else if indexCol >= 6 {
					boxes[8] = append(boxes[8], i)
				}
			}

			if bytes.Count(line, []byte{i}) > 1 {
				return false
			}
		}
	}
	for x := 0; x < 9; x++ { // Проверка всех столбцов на повторяющиеся символы
		mapCounter := make(map[byte]int)
		for i := 0; i < 9; i++ {
			if board[i][x] == '.' {
				continue
			}
			ok := mapCounter[board[i][x]]
			if ok == 0 {
				mapCounter[board[i][x]]++
				continue
			}
			return false

		}
	}

	for _, box := range boxes {
		for _, x := range box {
			if bytes.Count(box, []byte{x}) > 1 {
				return false
			}
		}
	}

	return true
}

// func checkBox(box []byte) bool {

// }
