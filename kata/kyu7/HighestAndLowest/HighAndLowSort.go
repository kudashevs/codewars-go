package highestandlowest

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLowSort(in string) string {
	chars := strings.Split(in, ` `)
	digits := make([]int, len(chars))

	for i, v := range chars {
		digits[i], _ = strconv.Atoi(v)
	}

	sort.Ints(digits)

	return fmt.Sprintf("%d %d", digits[len(digits)-1], digits[0])
}
