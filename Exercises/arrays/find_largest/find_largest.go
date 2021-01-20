package findlargest

// FindLargest functions
func FindLargest(a []int) (index int) {
	largest := 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[largest] {
			largest = i
		}
	}
	return largest
}
