package array

import (
	"fmt"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	number := findRepeatNumber(nums)
	fmt.Printf("number is %d", number)
}
