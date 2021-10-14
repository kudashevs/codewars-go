package sumofpositive

type positives []int

func PositiveSum(numbers []int) int {
	return filter(numbers, func(v int) bool { return (v >= 0) }).sum()
}

func filter(numbers []int, f func(v int) bool) (result positives) {
	for _, v := range numbers {
		if f(v) {
			result = append(result, v)
		}
	}

	return
}

func (data positives) sum() (result int) {
	for _, v := range data {
		result += v
	}

	return
}
