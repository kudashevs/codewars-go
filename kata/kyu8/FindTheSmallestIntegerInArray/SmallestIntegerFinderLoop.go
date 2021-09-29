package findsmallestinteger

func SmallestIntegerFinderLoop(numbers []int) int {
	min, numbers := numbers[0], numbers[1:]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}

	return min
}
