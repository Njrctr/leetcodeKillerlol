package leetcode

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	byteStr := []byte(s)
	var result [][]byte
	for x := numRows; x > 0; x-- {
		result = append(result, []byte{})
	}
	current := 0
	var forward bool = true
	for _, l := range byteStr {
		if current < 0 {
			forward = true
			current += 2
		} else if current == numRows {
			current -= 2
			forward = false
		}

		result[current] = append(result[current], l)
		if forward {
			current++
		} else {
			current--
		}

	}
	var correctResult []byte
	for _, x := range result {
		correctResult = append(correctResult, x...)
	}
	return string(correctResult)
}
