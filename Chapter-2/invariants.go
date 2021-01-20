package chapter2

func mininum(n int, a []float64) float64 {
	if n <= 0 { // preserve the case that n is smaller than starting point 1
		return 0
	}
	min := a[0]
	// min equals the minimum item from a[0],...,a[i-1]; i-1 == 0
	for i := 1; i != n; i++ {
		if a[i] < min { // min equals the minimum item in a[0],...,a[i-1]
			min = a[i]
		}
	}
	// min equals the minimum item from a[0],...,a[i-1]; i == n
	return min
}
