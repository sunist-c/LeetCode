package find_the_losers_of_the_circular_game

func circularGameLosers(n int, k int) []int {
	set := map[int]struct{}{}
	index, step := 1, 1
	for _, exist := set[index]; !exist; _, exist = set[index] {
		set[index] = struct{}{}
		index = reset(index+step*k, n)
		step++
	}

	var result []int
	for i := 1; i <= n; i++ {
		if _, ok := set[i]; !ok {
			result = append(result, i)
		}
	}

	return result
}

func reset(n, k int) int {
	for n > k {
		n -= k
	}
	return n
}
