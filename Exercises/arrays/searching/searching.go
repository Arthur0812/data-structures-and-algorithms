package searching

import (
	"math"
)

// LinearSearch searches x in the given array of numbers, using the linear searching technique
func LinearSearch(arr []int, x int) (index int) {
	n := len(arr)
	pos, right := -1, n-1
	for left := 0; left <= right; left, right = left+1, right-1 {
		if arr[left] == x {
			pos = left
			break
		}
		if arr[right] == x {
			pos = right
			break
		}
	}
	return pos
}

// BinarySearch searches x in the SORTED array of numbers, using the binary searching technique
func BinarySearch(arr []int, x int) (index int) {
	// ITEMS HAVE TO BE ORDERED
	n := len(arr)
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] < x { // search the right part
			left = mid + 1
		} else if arr[mid] == x { // match
			return mid
		} else { // search the left part
			right = mid - 1
		}
	}
	return -1
}

// JumpSearch searches x in the SORTED array of numbers, using the jump searching techinque
func JumpSearch(arr []int, x int) (index int) {
	// ITEMS HAVE TO BE ORDERED
	n := len(arr)
	step := int(math.Round(math.Sqrt(float64(n))))
	// find the position where x could possibly be located (between prev and curr)
	curr, prev := step, 0
	for arr[min(curr, n)-1] < x {
		prev, curr = curr, curr+step
		if prev >= n { // not found
			return -1
		}
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
