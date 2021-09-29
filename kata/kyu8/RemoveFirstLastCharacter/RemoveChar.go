package removefirstlastchar

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}
