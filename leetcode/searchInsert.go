package leetcode

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	if target < nums[left] {
		return 0
	}

	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left = mid + 1
			continue
		}
		if target < nums[mid] {
			right = mid - 1
			continue
		}
		return mid
	}
	return left
}
