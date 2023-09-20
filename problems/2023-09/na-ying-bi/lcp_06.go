package na_ying_bi

func minCount(coins []int) (result int) {
	for _, coin := range coins {
		if coin%2 == 0 {
			result += coin / 2
		} else {
			result += coin/2 + 1
		}
	}

	return result
}
