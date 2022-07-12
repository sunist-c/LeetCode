package day1

import "math"

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func PrimeCount(n int) int {
	// init
	ans := 0
	arr := make([]bool, n+1, n+1)
	arr[0], arr[1] = true, true

	// eratosthenes
	for i := 2; i <= n; i++ {
		// already judged
		if arr[i] {
			continue
		}

		// judge new number
		if isPrime(i) {
			ans += 1
			for j := 1; j*i <= n; j++ {
				arr[j*i] = true
			}
		}
	}

	return ans
}
