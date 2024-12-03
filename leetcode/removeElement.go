package leetcode

import "fmt"

func removeElement(nums []int, val int) int {
	counter := 0
	for x := 0; x < len(nums); x++ {
		if nums[x] == val {
			nums = append(nums[:x], nums[x+1:]...)
			x--
		} else {
			counter++
		}
	}
	fmt.Print(nums)
	return counter
}
