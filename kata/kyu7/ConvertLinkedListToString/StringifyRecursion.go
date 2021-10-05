package convertlinkedlisttostring

import "fmt"

func StringifyRecursion(list *Node) string {
	if list == nil {
		return `nil`
	}

	return fmt.Sprintf("%d", list.Data()) + ` -> ` + StringifyRecursion(list.Next())
}
