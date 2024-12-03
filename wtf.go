package main

import (
	"fmt"
)

//Given two binary strings a and b, return their sum as a binary string.

// Example 1:

// Input: a = "11", b = "1"
// Output: "100"

func main() {
	//a, b := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	// i, _ := strconv.ParseInt(a, 2, 64)
	a, b := "101111", "10" // "1000"

	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	ai := []byte(a)
	bi := []byte(b)
	increment := 0
	minlen := len(bi)
	longest := ai
	minimal := bi
	if len(ai) < len(bi) {
		minlen = len(ai)
		longest = bi
		minimal = ai
	}
	var last int
	for x := 1; x <= minlen; x++ {
		currA := &longest[len(longest)-x]
		currB := &minimal[len(minimal)-x]
		last = len(longest) - x - 1
		if *currA == 48 { //currA==48
			if increment == 0 { //currA==48, increment==0
				if *currB == 48 { //currA==48, increment==0, currB==48
					continue
				}
				//currA==48, increment==0, currB==49
				*currA++
				continue
			}
			//currA==48, increment!=0
			increment--
			*currA++
			if *currB == 48 { //currA==48, increment!=0, currB==48
				continue
			}
			//currA==48, increment!=0, currB==49
			increment++
			*currA--
		} else { //currA==49
			if increment == 0 { //currA==49, increment==0
				if *currB == 48 { //currA==49, increment==0, currB==48
					continue
				}
				//currA==49, increment==0, currB==49
				*currA--
				increment++
				continue
			}
			//currA==49, increment!=0
			*currA--
			if *currB == 48 { //currA==49, increment!=0, currB==48
				continue
			}
			//currA==49, increment!=0, currB==49
			*currA++
		}
	}
	fmt.Println("last", last, "increment", increment, "longest:", longest)
	for increment > 0 {
		if last < 0 {
			longest = append([]byte{49}, longest...)
			return string(longest)
		}
		if longest[last] == 48 { //longest[last]==48
			longest[last]++
			increment--
		} else { //longest[last]==49
			longest[last]--
			increment++
		}
		last--
	}

	return string(longest)
}
