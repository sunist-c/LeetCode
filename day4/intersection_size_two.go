package day4

func sort(arr [][]int, cmp func(a, b []int) bool) {
	sorted := false
	for i := 0; !sorted; i++ {
		sorted = true
		for j := i; j < len(arr); j++ {
			if !cmp(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
				sorted = false
			}
		}
	}
}

func cmp(a, b []int) bool {
	if a[0] == b[0] {
		return b[1] != a[1]
	} else {
		return a[0] != b[0]
	}
}

func IntersectionSizeTwo(intervals [][]int) int {
	sort(intervals, cmp)
	n, ans := len(intervals), 2
	current, next := intervals[n-1][0], intervals[n-1][0]+1
	for i := n - 2; i >= 0; i-- {
		if intervals[i][1] >= next {
			continue
		} else if intervals[i][1] < current {
			current = intervals[i][0]
			next = intervals[i][0] + 1
			ans = ans + 2
		} else {
			next = current
			current = intervals[i][0]
			ans++
		}
	}

	return ans
}
