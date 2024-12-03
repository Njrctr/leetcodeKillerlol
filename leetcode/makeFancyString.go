package leetcode

import "fmt"

func makeFancyString(s string) string {
	str := []rune(s)
	needdeleate := []int{}
	var lastletter rune = str[0]
	countrepeat := 0
	for x := 1; x < len(str); x++ {
		if str[x] == lastletter {
			countrepeat++
		} else {
			lastletter = str[x]
			countrepeat = 0
		}
		if countrepeat >= 2 {
			needdeleate = append(needdeleate, x)
		}
	}
	fmt.Println(needdeleate)
	for f := len(needdeleate) - 1; f >= 0; f-- {
		fmt.Println("f: ", f, "; str now: ", str)
		str = append(str[:needdeleate[f]], str[needdeleate[f]+1:]...)
	}
	return string(str)
}
func MakeFancyString(s string) string {
	chars := []rune(s)
	if len(chars) < 3 {
		return s
	}

	j := 2
	for i := 2; i < len(chars); i++ {
		fmt.Println("i: ", i, "\nj: ", j)
		fmt.Println("chars[i]: ", chars[i], "\nchars[j-1]: ", chars[j-1], "\nchars[j-2]: ", chars[j-2])
		if chars[i] != chars[j-1] || chars[i] != chars[j-2] {
			chars[j] = chars[i]
			j++
		}
	}

	return string(chars[:j])
}
