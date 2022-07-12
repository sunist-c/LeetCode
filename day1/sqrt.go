package day1

func SqrtWithBinary(x int) int {
	low, up, mid := 0, x, x
	for up != low {
		if up*up > x {
			up = low + (up-low)/2
		} else {
			low = up
			up = up + (mid-up)/2
		}
	}
	return up
}

func SqrtWithNewton(x int) int {
	if x == 0 {
		return 0
	}

	return int(newtonSqrt(float64(x), x))
}

func newtonSqrt(i float64, x int) float64 {
	res := (i + float64(x)/i) / 2
	if res == i {
		return i
	} else {
		return newtonSqrt(res, x)
	}
}
