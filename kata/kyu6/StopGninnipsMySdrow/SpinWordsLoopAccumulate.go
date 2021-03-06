package stopgninnipsmysdrow

import "strings"

func SpinWordsLoopAccumulate(str string) string {
	res := ``
	words := strings.Split(str, ` `)

	for _, v := range words {
		if len(v) > 4 {
			res += reverseAccumulate(v)
		} else {
			res += v
		}
		res += ` `
	}

	return strings.TrimSpace(res)
}

func reverseAccumulate(str string) (res string) {
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	return
}
