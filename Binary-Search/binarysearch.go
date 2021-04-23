//BinarySearch algorithm implementation
//Time complexity is O(logn)
package binarysearch

import (
	"fmt"
)

func BinarySearch(container []int, target int) int {
	start, end := 0, len(container) - 1

	for start < end {
		median := start + (end - start)/2
		if container[median] > target {
			end = median - 1
		} else if container[median] < target {
			start = median + 1
		} else {
			return median
		}
	}
	return -1
}
