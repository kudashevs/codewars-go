package vowelcount

func GetCountIntersection(str string) (count int) {
	intersect := intersection([]rune(`aeiou`), []rune(str))

	return len(intersect)
}

func intersection(a, b []rune) (c []rune) {
	hash := make(map[rune]bool)

	for _, item := range a {
		hash[item] = true
	}

	for _, item := range b {
		if hash[item] {
			c = append(c, item)
		}
	}

	return
}
