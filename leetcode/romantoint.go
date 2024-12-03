package leetcode

import (
	"strings"
)

func romanToInt(s string) int {
	slc := strings.Split(s, "")
	RomanNums := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	totalSumm := 0
	for i, l := range slc {
		if i == len(slc)-1 {
			totalSumm = totalSumm + RomanNums[l]
			break
		}
		if RomanNums[l] < RomanNums[slc[i+1]] {
			totalSumm = totalSumm - RomanNums[l]
		} else {
			totalSumm = totalSumm + RomanNums[l]
		}
	}
	return totalSumm
}
