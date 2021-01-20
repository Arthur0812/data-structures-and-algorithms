package fibonacci

// Fibonacci returns the n-numbers long fibonacci sequence
func Fibonacci(n int) []int {
	fibonacci := make([]int, n, n)
	for i := range fibonacci {
		if i == 0 || i == 1 { // base cases
			fibonacci[i] = 1
			continue
		}
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	return fibonacci
}
