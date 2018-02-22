package sort

import (
	"reflect"
	"testing"
)

func TestSortTwoSlicesInts(t *testing.T) {
	tests := []struct {
		firstSlice         []int
		secondSlice        []int
		updatedFirstSlice  []int
		updatedSecondSlice []int
	}{
		{
			firstSlice:         []int{1, 2, 3},
			secondSlice:        []int{1, 0, 0},
			updatedFirstSlice:  []int{2, 3, 1},
			updatedSecondSlice: []int{0, 0, 1},
		},
		{
			firstSlice:         []int{1, 2, 3},
			secondSlice:        []int{0, 1, 0},
			updatedFirstSlice:  []int{1, 3, 2},
			updatedSecondSlice: []int{0, 0, 1},
		},
	}

	for _, test := range tests {
		SortTwoSlicesInts(test.firstSlice, test.secondSlice)
		if !reflect.DeepEqual(test.firstSlice, test.updatedFirstSlice) {
			t.Errorf("[TestSortTwoSlicesInts] got first slice: %v, want first slice: %v\n", test.firstSlice, test.updatedFirstSlice)
		}
		if !reflect.DeepEqual(test.secondSlice, test.updatedSecondSlice) {
			t.Errorf("[TestSortTwoSlicesInts] got second slice: %v, want second slice: %v\n", test.secondSlice, test.updatedSecondSlice)
		}
	}
}
