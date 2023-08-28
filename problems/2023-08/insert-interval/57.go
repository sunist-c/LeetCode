package insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var result [][]int
	scopeStart, scopeEnd := newInterval[0], newInterval[1]
	for _, item := range intervals {
		start, end := item[0], item[1]
		if end < scopeStart {
			result = append(result, []int{start, end})
		} else if start > scopeEnd {
			result = append(result, []int{scopeStart, scopeEnd})
			scopeStart, scopeEnd = start, end
		} else {
			scopeStart, scopeEnd = min(scopeStart, start), max(scopeEnd, end)
		}
	}
	result = append(result, []int{scopeStart, scopeEnd})
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}
