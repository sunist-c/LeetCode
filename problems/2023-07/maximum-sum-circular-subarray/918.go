package maximum_sum_circular_subarray

import "math"

func maxSubarraySumCircular(nums []int) int {
	maxSubArraySum := math.MinInt
	minSubArraySum := 0
	var maxF, minF, sum int
	for _, x := range nums {
		maxF = max(maxF, 0) + x
		maxSubArraySum = max(maxSubArraySum, maxF)
		minF = min(minF, 0) + x
		minSubArraySum = min(minSubArraySum, minF)
		sum += x
	}
	if sum == minSubArraySum {
		return maxSubArraySum
	}
	return max(maxSubArraySum, sum-minSubArraySum)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
