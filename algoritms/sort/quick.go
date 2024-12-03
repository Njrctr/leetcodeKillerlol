package sort

func Quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	var less, greater []int

	for _, num := range arr[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(Quick(less), pivot)
	result = append(result, Quick(greater)...)
	return result
}
