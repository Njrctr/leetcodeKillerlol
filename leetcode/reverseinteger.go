package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

func _() {
	x := 8
	fmt.Println(myreverse(x))
}

func myreverse(x int) int {
	if x > 10 && x < 10 {
		return x
	}
	stringInt := strconv.Itoa(x)
	strRunes := []rune(stringInt)
	reversedRunes := []rune{}
	if strRunes[0] == []rune("-")[0] {
		reversedRunes = append(reversedRunes, strRunes[0])
		strRunes = strRunes[1:]
	}
	for x := len(strRunes) - 1; x >= 0; x-- {
		reversedRunes = append(reversedRunes, strRunes[x])
	}
	reversedInt, _ := strconv.Atoi(string(reversedRunes))
	if reversedInt < math.MinInt32 || reversedInt > math.MaxInt32 {
		return 0
	}
	return reversedInt
}

func bestreverse(x int) int {
	x32 := int32(x)

	result := int32(0)
	base := int32(10)
	for x32 != 0 {
		d := x32 % base
		x32 /= base

		if (result-d)/base != result {
			return 0
		} // detected an overflow
		newResult := result*base + d

		result = newResult
	}

	return int(result)
}
