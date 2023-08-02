package card_flipping_game

import (
	"math"
)

func flipgame(fronts []int, backs []int) int {
	cannot := map[int]struct{}{}
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			cannot[fronts[i]] = struct{}{}
		}
	}

	result := math.MaxInt
	for i := 0; i < len(fronts); i++ {
		if _, exist := cannot[fronts[i]]; !exist {
			result = min(result, fronts[i])
		}
		if _, exist := cannot[backs[i]]; !exist {
			result = min(result, backs[i])
		}
	}

	if result == math.MaxInt {
		result = 0
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
