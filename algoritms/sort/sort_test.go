package sort

import (
	"alg/src"
	"alg/tracking"
	"sort"
	"testing"
)

var unsortedSlice []int = src.CreateArr(10000)

func TestBubbleSort(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения Пузырьковой сортировки: "))
	Bubble(unsortedSlice)
}

func TestSelectionSort(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения сортировки Выбором: "))
	Selection(unsortedSlice)
}

func TestInsertionSort(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения сортировки Вставкой: "))
	Insertion(unsortedSlice)
}

func TestQuickSort(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения Быстрой сортировки: "))
	Quick(unsortedSlice)
}

func TestMergeSort(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения сортировки Слиянием: "))
	Merge(unsortedSlice)
}

func TestStandartGoSort(t *testing.T) {
	defer tracking.Duration(tracking.Track("Время выполнения Стандартной Go сортировки: "))
	sort.Ints(unsortedSlice)
}
