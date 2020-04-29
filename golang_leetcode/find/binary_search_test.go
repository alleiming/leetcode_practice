package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	len := 1000
	array := make([]int, len)

	for i := 0; i < len; i++ {
		array[i] = i
	}

	fmt.Print(binarySearch(array, 66))
}
