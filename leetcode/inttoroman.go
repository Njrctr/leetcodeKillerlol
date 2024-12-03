package leetcode

import (
	"fmt"
	"strings"
)

func _() {
	s := 58 // MMMDCCXLIX
	fmt.Println(intToRoman(s))
}

func intToRoman(num int) string {
	counter := 1
	result := []string{}
	for num > 0 {
		newres := []string{getroman(num%10, counter)}
		newres = append(newres, result...)
		num = num / 10
		counter *= 10
		result = newres

	}
	// fmt.Println("result: ", result)
	return strings.Join(result, "")
}

func getroman(num, c int) string {
	RomanNums := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}
	if num == 4 {
		return RomanNums[1*c] + RomanNums[5*c]
	} else if num == 9 {
		return RomanNums[1*c] + RomanNums[10*c]
	} else if num >= 5 {
		return RomanNums[5*c] + strings.Repeat(RomanNums[1*c], num-5)
	}
	return strings.Repeat(RomanNums[1*c], num)
}
