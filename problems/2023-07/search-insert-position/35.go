package search_insert_position

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, len(nums), 0)
}

func binarySearch(nums []int, target int, upper, lower int) int {
	if upper == lower {
		return upper
	}

	mid := (upper + lower) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(nums, target, mid, lower)
	} else {
		return binarySearch(nums, target, upper, mid+1)
	}
}
