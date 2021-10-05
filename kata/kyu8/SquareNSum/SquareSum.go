package squarensum

func SquareSum(numbers []int) (res int) {
	for _, v := range numbers {
		res += v * v
	}

	return
}
