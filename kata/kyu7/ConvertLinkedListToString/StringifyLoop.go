package convertlinkedlisttostring

import "fmt"

func StringifyLoop(list *Node) (res string) {
	for list != nil {
		res += fmt.Sprintf("%d", list.Data()) + ` -> `
		list = list.Next()
	}

	res += `nil`

	return
}
