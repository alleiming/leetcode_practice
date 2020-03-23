package main

import "fmt"

//两数之和
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{4, 93, 84, 85, 80, 37, 81, 93, 27, 12}

	result := twoSum(nums, 39)

	fmt.Print(result)
}
