package day6

import (
	"sort"
)

func ArrayRankTransform(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] < tmp[j]
	})
	m := map[int]int{}

	index := 1
	m[tmp[0]] = index
	for i := 1; i < len(tmp); i++ {
		if tmp[i-1] == tmp[i] {
			continue
		} else {
			index += 1
			m[tmp[i]] = index
		}
	}

	for i, v := range arr {
		tmp[i] = m[v]
	}
	return tmp
}
