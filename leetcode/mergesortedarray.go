package leetcode

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	// if n == 0 {
	// 	return nums1
	// }
	for j, i := 0, m; j < n; j, i = j+1, i+1 {
		nums1[i] = nums2[j]
	}
	sort.Ints(nums1)
	return nums1
}
