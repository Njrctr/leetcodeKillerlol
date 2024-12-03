package sort

func Merge(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := Merge(arr[:mid])
	right := Merge(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] { // Сравниваем элементы из левой и правой половины
			merged = append(merged, left[0]) // Если из левой меньше - добавляем его
			left = left[1:]
		} else {
			merged = append(merged, right[0]) // иначе добавляем из правой
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}
