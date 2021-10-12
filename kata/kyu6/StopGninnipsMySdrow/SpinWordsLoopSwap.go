package stopgninnipsmysdrow

import "strings"

func SpinWordsLoopSwap(str string) string {
	res := ``
	words := strings.Split(str, ` `)

	for _, v := range words {
		if len(v) > 4 {
			res += reverseSwap(v)
		} else {
			res += v
		}
		res += ` `
	}

	return strings.TrimSpace(res)
}

func reverseSwap(str string) string {
	bytes := []rune(str)

	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)
}
