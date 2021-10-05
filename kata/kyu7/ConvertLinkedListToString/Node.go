package convertlinkedlisttostring

type Node struct {
	data int
	next *Node
}

func NewNode(data int, next *Node) *Node {
	return &Node{data: data, next: next}
}

func (n *Node) Data() int {
	return n.data
}

func (n *Node) Next() *Node {
	return n.next
}
