package toeplitz_matrix

func isToeplitzMatrix(matrix [][]int) bool {
	height, width := len(matrix), len(matrix[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i > 0 && j > 0 && matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}
