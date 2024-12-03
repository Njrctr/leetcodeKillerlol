package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	left := 0
	if len(nums) == 0 {
		return 0
	}
	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			fmt.Print("\nOld nums[left]: ", nums[left], ". nums[right]: ", nums[right])
			nums[left+1] = nums[right]
			left++
		}
	}
	fmt.Print("unique: ", len(nums[:left+1]), "\n")
	return len(nums[:left+1])
}
