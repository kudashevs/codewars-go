package highestandlowest

import (
	"strconv"
	"strings"
)

func HighAndLowLoop(in string) string {
	chars := strings.Split(in, ` `)
	digits := make([]int, len(chars))

	for i, v := range chars {
		digits[i], _ = strconv.Atoi(v)
	}

	return findMaxMin(digits)
}

func findMaxMin(data []int) string {
	var min int = data[0]
	var max int = data[0]

	for _, v := range data {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return string(strconv.Itoa(max) + ` ` + strconv.Itoa(min))
}
