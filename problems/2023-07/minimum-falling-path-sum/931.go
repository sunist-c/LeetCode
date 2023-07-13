package minimum_falling_path_sum

import "math"

func minFallingPathSum(matrix [][]int) int {
	length := len(matrix)
	var dp [][]int
	for i := 0; i < length; i++ {
		// 多开辟两个空间避免越界
		dp = append(dp, make([]int, length+2))
		dp[i][0], dp[i][length+1] = math.MaxInt, math.MaxInt
	}

	// 初始化第一行
	for i := 0; i < length; i++ {
		dp[0][i+1] = matrix[0][i]
	}

	// 从第二行开始进行动态规划
	for i := 1; i < length; i++ {
		for columns, number := range matrix[i] {
			dp[i][columns+1] = min(
				dp[i-1][columns],
				dp[i-1][columns+1],
				dp[i-1][columns+2],
			) + number
		}
	}

	return min(dp[length-1]...)
}

func min(numbers ...int) int {
	result := math.MaxInt
	for _, number := range numbers {
		if number < result {
			result = number
		}
	}
	return result
}
