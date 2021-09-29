package findsmallestinteger

import "sort"

func SmallestIntegerFinderSort(numbers []int) int {
	sort.Ints(numbers)

	return numbers[0]
}
