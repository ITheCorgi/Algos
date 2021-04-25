//Merge sort implementation.
//Time complexity is O(logn)
package mergesort

import (
	"fmt"
)

func MergeSort(in []int) []int {
	if len(in) < 2 {
		return in
	}
	mid := len(in)/2
	return merge(mergeSort(in[:mid]), mergeSort(in[mid:]))
}

func merge(left, right []int) []int {
	newSize, lidx, ridx := len(left) + len(right), 0, 0
	union := make([]int, newSize)
	
	for i:=0; i<newSize; i++ {
		switch {
		//in case we have already passed through left part, just copy right to result
		case left[lidx] <= lidx:
			union[i] = right[ridx]
			ridx++
		//in case we have already passed through right part, just copy left to result
		case right[ridx] <= ridx:
			union[i] = left[lidx]
			lidx++

		case left[lidx] < right[ridx]:
			union[i] = left[lidx]
			lidx++

		default:
			union[i] = right[ridx]
			ridx++			
		}
	}
	return union
} 