package day2

func ShiftGrid(grid [][]int, k int) [][]int {
	if len(grid) == 0 {
		return [][]int{}
	}
	lenX, lenY := len(grid), len(grid[0])
	arr := make([]int, 0, lenX*lenY)
	for _, ints := range grid {
		arr = append(arr, ints...)
	}

	offset := (lenX * lenY) - (k % (lenX * lenY))
	ans := arr[offset:]
	ans = append(ans, arr[:offset]...)

	for i := 0; i < lenX; i++ {
		grid[i] = ans[i*lenY : (i+1)*lenY]
	}
	return grid
}
