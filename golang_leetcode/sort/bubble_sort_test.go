package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 3, 2, 7, 6, 7, 9, 1, 10}
	bubbleSort(arr)
	for _, item := range arr {
		fmt.Printf("%d ", item)
	}
}
