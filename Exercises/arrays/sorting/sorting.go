package sorting

// BubbleSort sorts the given array of numbers using the bubble sorting technique
func BubbleSort(arr []int) {
	n := len(arr)

	for i := n - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// SelectionSortInt sorts the given array of numbers using the selection sorting technique
func SelectionSortInt(arr []int) {
	n := len(arr)
	// loop from start to the penultimate element
	for i := 0; i < n-1; i++ {
		// the index of the smallest element starting from i+1 up to n
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// a smaller element was found
		if minIndex != i {
			// swap the smallest found element (from i+1 to n) with the current main element (at i)
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
}

// SelectionSortString sorts the given array of strings using the selection sorting technique
func SelectionSortString(arr []string) {
	n := len(arr)
	// loop from start to penultimate element
	for i := 0; i < n-1; i++ {
		// the index of the smallest string startting from i+1 up to n
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// a smaller string was found
		if minIndex != i {
			// swap the smallest found string
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

// InsertionSort sorts a given array of numbers using the insertion sorting technique
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		curr, j := arr[i], i-1
		for ; j >= 0 && curr <= arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = curr
	}
}

// MergeSort sorts the given array of numbers using the merge sorting technique
func MergeSort(arr []int) {
	n := len(arr)
	for currSize := 1; currSize < n; currSize *= 2 {
		for left := 0; left < n-1; left += 2 * currSize {
			mid := min(left+currSize-1, n-1)
			right := min(left+2*currSize-1, n-1)

			merge(arr, left, mid, right)
		}
	}
}

func merge(arr []int, left, mid, right int) {
	leftN := mid - left + 1
	rightN := right - mid

	// temporary slices of left and right side
	leftS := make([]int, leftN)
	rightS := make([]int, rightN)

	// fill the left and right slices with corresponding elements from arr
	for i := 0; i < leftN; i++ {
		leftS[i] = arr[i+left]
	}
	for j := 0; j < rightN; j++ {
		rightS[j] = arr[j+mid+1]
	}

	i, j := 0, 0
	k := left

	// merge the left and right onto arr (small to larger, ascending order)
	for i < leftN && j < rightN {
		if leftS[i] <= rightS[j] {
			arr[k] = leftS[i]
			i++
		} else {
			arr[k] = rightS[j]
			j++
		}
		k++
	}
	// add remaining elements from either left or right
	for i < leftN {
		arr[k] = leftS[i]
		i, k = i+1, k+1
	}
	for j < rightN {
		arr[k] = rightS[j]
		j, k = j+1, k+1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// HeapSort sorts the given array of numbers using the heap sorting technique
func HeapSort(arr []int) {
	n := len(arr)

	// rearrange the array to be "heapifieable"
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// one by one, sort the array
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapify(arr, i, 0)
	}
}

func heapify(arr []int, size int, index int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < size && arr[left] > arr[largest] { // left is greatest
		largest = left
	}
	if right < size && arr[right] > arr[largest] { // right is greatest
		largest = right
	}

	// if either the left child or the right child was greater, swap it with the parent
	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]

		// recursively heapify the sub-tree (either the right or left child)
		heapify(arr, size, largest)
	}
}
