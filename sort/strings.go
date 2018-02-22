package sort

import "strings"

func SortTwoSlices(firstSlice, secondSlice []string) {
	for i := 0; i < len(firstSlice)-1; i++ {
		min := i
		for j := i + 1; j < len(firstSlice); j++ {
			if strings.Compare(secondSlice[j], secondSlice[min]) == -1 {
				min = j
			}
		}

		tempFirst := firstSlice[min]
		firstSlice[min] = firstSlice[i]
		firstSlice[i] = tempFirst

		tempSecond := secondSlice[min]
		secondSlice[min] = secondSlice[i]
		secondSlice[i] = tempSecond
	}
}
