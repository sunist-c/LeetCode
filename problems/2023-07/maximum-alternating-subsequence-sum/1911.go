package maximum_alternating_subsequence_sum

func maxAlternatingSum(nums []int) int64 {
	x, y := 0, 0
	for _, i := range nums {
		x, y = max(x, y+i), max(y, x-i)
	}
	return int64(max(x, y))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
