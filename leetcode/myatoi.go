package leetcode

import (
	"fmt"
	"math"
	"strings"
)

func _() {
	str := "-91283472332"
	fmt.Println(myAtoi(str))
}

func myAtoi(s string) int {
	str := strings.TrimSpace(s)
	if len(str) == 0 {
		return 0
	}
	var negative bool
	if str[0] == '-' {
		negative = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	var res int
	for i, cur := range str {
		if cur < '0' || cur > '9' {
			break
		}
		res = res*10 + int(str[i]-'0')

		if res > math.MaxInt32 && !negative {
			return math.MaxInt32
		}
		if -res < math.MinInt32 && negative {
			return math.MinInt32
		}
	}
	if negative {
		res = -res
	}
	return res
}
