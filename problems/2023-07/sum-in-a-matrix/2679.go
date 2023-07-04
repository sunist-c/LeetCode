package sum_in_a_matrix

import (
	"sort"
)

func matrixSum(nums [][]int) int {
	height, length := len(nums), len(nums[0])
	for i, array := range nums {
		sort.Slice(array, func(i, j int) bool {
			return array[i] > array[j]
		})
		nums[i] = array
	}

	result := 0
	for i := 0; i < length; i++ {
		maxNumber := 0
		for j := 0; j < height; j++ {
			if maxNumber < nums[j][i] {
				maxNumber = nums[j][i]
			}
		}
		result += maxNumber
	}

	return result
}
