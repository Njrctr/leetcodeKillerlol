package leetcode

import (
	"strconv"
	"strings"
)

func hammingWeight(n uint32) int {
	res := 0
	for n != 0 {
		if n&1 == 1 {
			res++

		}
		n >>= 1
	}
	return res
}

func hammingWeight2(n uint32) int {
	binaryRepresentation := strconv.FormatInt(int64(n), 2)
	return strings.Count(binaryRepresentation, "1")
}
