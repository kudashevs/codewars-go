package splitstrings

func SplitStringSplit(str string) []string {
	return splitByNWithoutRemainder(str+`_`, 2)
}

func splitByNWithoutRemainder(str string, n int) (res []string) {
	length := len(str)

	for i := 0; i < length; i++ {
		if i != 0 && i%n == 0 {
			res = append(res, str[i-n:i])
		}
	}

	if length%n == 0 {
		res = append(res, str[length-n:])
	}

	return
}
