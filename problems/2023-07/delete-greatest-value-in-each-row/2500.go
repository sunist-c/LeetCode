package delete_greatest_value_in_each_row

import "sort"

func deleteGreatestValue(grid [][]int) int {
	for _, line := range grid {
		sort.Ints(line)
	}

	result := 0
	for i := 0; i < len(grid[0]); i++ {
		result += maxColNum(grid, i)
	}

	return result
}

func maxColNum(grid [][]int, col int) (max int) {
	for i := 0; i < len(grid); i++ {
		if grid[i][col] > max {
			max = grid[i][col]
		}
	}
	return max
}
