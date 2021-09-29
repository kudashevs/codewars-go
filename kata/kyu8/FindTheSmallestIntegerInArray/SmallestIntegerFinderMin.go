package findsmallestinteger

func SmallestIntegerFinderMin(numbers []int) int {
	return min(numbers)
}

func min(arr []int) int {
	min, arr := arr[0], arr[1:]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}
