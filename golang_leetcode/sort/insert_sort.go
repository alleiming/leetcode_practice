package sort

//插入排序
func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			} else {
				break
			}
		}
	}
}

//改进后的插入排序，减少元素交换
func insertSortImprove(arr []int) {
	for i := 1; i < len(arr); i++ {
		e := arr[i] //保存需要交换的元素
		var j int
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = e
	}
}
