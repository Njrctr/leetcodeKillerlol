package leetcode

// Find the Power of K-Size Subarrays I

// You are given an array of integers nums of length n and a positive integer k.

// The power of an array is defined as:

//     Its maximum element if all of its elements are consecutive and sorted in ascending order.
//     -1 otherwise.

// You need to find the power of all subarraysof nums of size k.

// Return an integer array results of size n - k + 1, where results[i] is the power of nums[i..(i + k - 1)].

// Example 1:

// Input: nums = [1,2,3,4,3,2,5], k = 3

// Output: [3,4,-1,-1,-1]

// Explanation:

// There are 5 subarrays of nums of size 3:

//     [1, 2, 3] with the maximum element 3.
//     [2, 3, 4] with the maximum element 4.
//     [3, 4, 3] whose elements are not consecutive.
//     [4, 3, 2] whose elements are not sorted.
//     [3, 2, 5] whose elements are not consecutive.

import (
	"fmt"
)

func _() {
	nums := []int{1, 2, 3, 4, 3, 2, 5}
	fmt.Println(resultsArray(nums, 3))
}

func resultsArray(nums []int, k int) []int {

	if len(nums) == k {
		if k == 1 {
			return []int{nums[k-1]}
		}
		// Ищем максимальный, если упорядочен список, иначе возвращаем -1

		if nums[0] >= nums[k-1] {
			return []int{-1}
		}
		maxel := nums[0]
		for el := 1; el < k; el++ {
			if nums[el]-maxel == 1 {
				maxel = nums[el]
			} else {
				return []int{-1}
			}
		}
		return []int{maxel}
	}
	countsubarray := len(nums) - k + 1
	end := k
	var resultSlice []int
	for x := 0; x < countsubarray; x++ {
		res := resultsArray(nums[x:end], k)
		resultSlice = append(resultSlice, res[0])
		end++
	}
	return resultSlice
}
