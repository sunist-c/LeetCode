package three_sum_closest

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if abs(target-sum) < abs(target-ans) {
				ans = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return ans
			}
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
