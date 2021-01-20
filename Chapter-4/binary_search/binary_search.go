package binarysearch

func binarySearch(a []int, x int) (index int) {
	left, mid, right := 0, 0, len(a)-1
	for left < right {
		mid = (left + right) / 2
		if x > a[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if a[left] == x {
		return left
	}
	return -1
}
