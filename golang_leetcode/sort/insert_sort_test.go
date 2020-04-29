package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{5, 3, 2, 7, 6, 7, 9, 1, 10}
	insertSort(arr)
	for _, item := range arr {
		fmt.Printf("%d ", item)
	}
}

func TestInsertSortImprove(t *testing.T) {
	arr := []int{5, 3, 2, 7, 6, 7, 9, 1, 10}
	insertSortImprove(arr)
	for _, item := range arr {
		fmt.Printf("%d ", item)
	}
}
