package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 2, 7, 6, 7, 9, 1, 10}
	selectionSort(arr)
	for _, item := range arr {
		fmt.Printf("%d ", item)
	}
}
