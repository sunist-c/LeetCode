package day1

func SortedArrayDuplicates(nums []int) int {
	// empty array
	if len(nums) == 0 {
		return 0
	}

	// init
	ans := 1
	slow, fast := 0, 1

	// duplicate
	for fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			nums[slow+1] = nums[fast]
			slow++
			fast++
			ans++
		}
	}

	return ans
}
