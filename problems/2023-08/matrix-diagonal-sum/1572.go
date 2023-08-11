package matrix_diagonal_sum

func diagonalSum(mat [][]int) int {
	sum := 0
	for line, subArr := range mat {
		sum += subArr[line]
	}

	for line, subArr := range mat {
		if len(subArr)-1-line == line {
			continue
		}
		sum += subArr[len(subArr)-1-line]
	}

	return sum
}
