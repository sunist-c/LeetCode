package count_servers_that_communicate

func countServers(grid [][]int) int {
	rowsCounter, colsCounter := make([]int, len(grid)), make([]int, len(grid[0]))
	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				rowsCounter[i]++
				colsCounter[j]++
			}
		}
	}

	result := 0
	for i, row := range grid {
		for j, col := range row {
			if col == 1 && (rowsCounter[i] > 1 || colsCounter[j] > 1) {
				result++
			}
		}
	}

	return result
}
