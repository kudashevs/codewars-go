package vowelcount

import "strings"

func GetCountFilter(str string) (count int) {
	return len(filter([]rune(str), func(val rune) bool {
		return strings.ContainsRune(`aeiou`, val)
	}))
}

func filter(arr []rune, cond func(rune) bool) []rune {
	res := []rune{}
	for i := range arr {
		if cond(arr[i]) {
			res = append(res, arr[i])
		}
	}

	return res
}
