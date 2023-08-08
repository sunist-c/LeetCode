package maximum_absolute_sum_of_any_subarray

func maxAbsoluteSum(nums []int) int {
	maxSum, minSum := 0, 0
	maxAns := 0
	for _, num := range nums {
		maxSum = max(maxSum+num, num)
		minSum = min(minSum+num, num)
		maxAns = max(maxAns, max(abs(maxSum), abs(minSum)))
	}

	return maxAns
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
