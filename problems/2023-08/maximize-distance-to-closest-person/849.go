package maximize_distance_to_closest_person

func maxDistToClosest(seats []int) int {
	maxDist := 0
	for i, seat := range seats {
		if seat == 1 {
			continue
		}

		prev, next := findSeat(seats, i)
		maxDist = max(maxDist, min(prev, next))
	}

	return maxDist
}

func findSeat(seats []int, now int) (prev, next int) {
	for i := now; i >= 0; i-- {
		if seats[i] == 1 {
			prev = now - i
			goto findNext
		}
	}
	prev = 1 << 15

findNext:
	for i := now; i < len(seats); i++ {
		if seats[i] == 1 {
			next = i - now
			goto returnResult
		}
	}
	next = 1 << 15

returnResult:
	return prev, next
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
