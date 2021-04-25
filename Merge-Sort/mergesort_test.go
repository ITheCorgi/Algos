package mergesort

import (
	"fmt"
	"testing"
	"reflect"
)

func MergeSortTest(t *testing.T) {
	testcases := []struct {
		Name     string
		Input    []int
		Expected []int
	}{
		{
			"Empty slice",
			[]int{},
			[]int{},
		},
		{
			"One element dimension",
			[]int{1},
			[]int{1},
		},
		{
			"Slice with numbers",
			[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"Slice with repeatable numbers",
			[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6, 1, 10, 2, 9, 3, 8, 4, 7, 5, 6, 1, 10, 2, 9, 3, 8, 4, 7, 5, 6, 1, 10, 2, 9, 3, 8, 4, 7, 5, 6, 1, 10, 2, 9, 3, 8, 4, 7, 5, 6},
			[]int{1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			var got []int
			got = MergeSort(testcase.Input)
			if !reflect.DeepEqual(got, testcase.Expected) {
				t.Errorf("Test \"%s\" failed.\nExpected: %v\n Got: %v", testcase.Name, testcase.Expected, got)
			}
		})
	}
}