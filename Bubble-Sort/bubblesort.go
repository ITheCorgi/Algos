//Bubble sort implementation
//Time complexity O(n^2), space complexity O(1)
package bubblesort

import (
	"fmt"
)

func bubbleSort(in []int) {
	var isSorted bool
	for isSorted == false {
		isSorted = true
		for i:=0; i<len(in); i++ {
			if in[i] > in[i+1] {
				in[i], in[i+1] = in[i+1], in[i]
				isSorted = false
			}
		}
	}
}

/*
func main() {
	arr := []int {10, 3, 5, 6, 3, 4, 22, 65, 333, 11, 2, 15, 14}
	bubbleSort(arr)
	fmt.Println(arr)
}
*/