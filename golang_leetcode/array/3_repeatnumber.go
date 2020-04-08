package array

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
//查出数据中重复的数字
func findRepeatNumber(nums []int) int {
	len := len(nums)
	for i := 0; i < len; i++ {
		for {
			if i != nums[i] {
				if nums[i] == nums[nums[i]] {
					return nums[i]
				}

				temp := nums[i]
				nums[i] = nums[temp]
				nums[temp] = temp
			} else {
				break
			}
		}
	}
	return -1
}
