//Insertion sort algo with time complexity O(n^2)
package insertionsort

import (
	"fmt"
)

func insertionSort(in []int) {
	for i := 1; i < len(in); i++ {
		j := i
		for j > 0 {
			if in[j-1] > in[j] {
				in[j-1], in[j] = in[j], in[j-1]
			}
			j--
		}
	}
}