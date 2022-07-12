package day1

func PivotIndex(nums []int) int {
	// empty array
	if len(nums) == 0 {
		return -1
	}

	// init
	ans := -1
	sum := 0

	// calculate
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
	}

	// find
	left, right := 0, sum
	for i := 0; i < len(nums); i++ {
		if left == right {
			return i
		}
		left += nums[i]
		if i+1 < len(nums) {
			right -= nums[i+1]
		}
	}

	return ans
}
