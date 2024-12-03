package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	var countOfFriends int
// 	var strs []string
// 	fmt.Fscan(reader, &countOfFriends)
// 	x := 0
// 	for x < countOfFriends {
// 		var now string
// 		fmt.Fscan(reader, &now)
// 		strs = append(strs, now)
// 		x++
// 	}
// 	checkStrings(strs)

// }

func checkStrings(strs []string) {
	uniqueCount := 0
	var lowely string
	for _, str := range strs {
		uniqletter := 0
		lenofword := len(str)
		strRune := []rune(str)
		for lenofword > 0 {
			lenofword -= strings.Count(str, string(strRune[lenofword-1]))
			uniqletter++
		}
		if uniqletter > uniqueCount {
			lowely = str
			uniqueCount = uniqletter
		}

	}
	fmt.Println(uniqueCount, lowely)
}
