package main

import "fmt"

//二分查找
func binarySearch(array []int, target int) int {
	var l, r = 0, len(array) - 1
	for {
		if l <= r {
			mid := (l + r) / 2
			if array[mid] == target {
				return mid
			}

			if array[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			break
		}
	}

	return -1
}

func main() {
	len := 1000
	array := make([]int, len)

	for i := 0; i < len; i++ {
		array[i] = i
	}

	fmt.Print(binarySearch(array, 999))
}
