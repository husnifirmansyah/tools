package sort

func SortTwoSlicesInts(firstSlice, secondSlice []int) {
	for i := 0; i < len(firstSlice)-1; i++ {
		min := i
		for j := i + 1; j < len(firstSlice); j++ {
			if secondSlice[j] < secondSlice[min] {
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
