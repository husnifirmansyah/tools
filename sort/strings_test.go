package sort

import (
	"reflect"
	"testing"
)

func TestSortTwoSlices(t *testing.T) {
	tests := []struct {
		firstSlice         []string
		secondSlice        []string
		updatedFirstSlice  []string
		updatedSecondSlice []string
	}{
		{
			firstSlice:         []string{"1", "2", "3"},
			secondSlice:        []string{"1", "0", "0"},
			updatedFirstSlice:  []string{"2", "3", "1"},
			updatedSecondSlice: []string{"0", "0", "1"},
		},
		{
			firstSlice:         []string{"1", "2", "3"},
			secondSlice:        []string{"0", "1", "0"},
			updatedFirstSlice:  []string{"1", "3", "2"},
			updatedSecondSlice: []string{"0", "0", "1"},
		},
	}

	for _, test := range tests {
		SortTwoSlices(test.firstSlice, test.secondSlice)
		if !reflect.DeepEqual(test.firstSlice, test.updatedFirstSlice) {
			t.Errorf("[TestSortTwoSlices] got first slice: %v, want first slice: %v\n", test.firstSlice, test.updatedFirstSlice)
		}
		if !reflect.DeepEqual(test.secondSlice, test.updatedSecondSlice) {
			t.Errorf("[TestSortTwoSlices] got second slice: %v, want second slice: %v\n", test.secondSlice, test.updatedSecondSlice)
		}
	}
}
