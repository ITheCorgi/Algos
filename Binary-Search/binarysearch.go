package binarysearch

import (
	"fmt"
)

//BinarySearch algorithm implementation
//Time complexity is O()
func BinarySearch(container []int, target int) int {
	start, end := 0, len(container) - 1

	for start < end {
		median := int(start + (end - start))/2
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