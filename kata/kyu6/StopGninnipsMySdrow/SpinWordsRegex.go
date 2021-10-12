package stopgninnipsmysdrow

import "regexp"

func SpinWordsRegex(str string) string {
	regex := regexp.MustCompile(`\w{5,}`)

	return regex.ReplaceAllStringFunc(str, func(w string) string { return reverse(w) })
}

func reverse(str string) (res string) {
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}

	return
}
