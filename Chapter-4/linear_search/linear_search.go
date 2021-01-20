package linearsearch

func linearSearch(a []int, x int) (index int) {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return i
		}
	}
	return -1
}
