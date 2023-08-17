package how_many_numbers_are_smaller_than_the_current_number

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	sort.Slice(copied, func(i, j int) bool {
		return copied[i] < copied[j]
	})

	counter := map[int]int{}
	for i, num := range copied {
		if _, exist := counter[num]; !exist {
			counter[num] = i
		}
	}

	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = counter[num]
	}

	return result
}
