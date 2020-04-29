package main

//二分查找
func binarySearch(array []int, target int) int {
	var left, right = 0, len(array) - 1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
