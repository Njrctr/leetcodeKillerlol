package src

import "math/rand"

func CreateArr(size int) (arr []int) {
	arr = make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func GetRange(end int) (result []int) {
	result = make([]int, 0, end)
	for value := 0; value <= end; value++ {
		result = append(result, value)
	}
	return result
}
