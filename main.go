package linkedList

// List is a thing
type List struct {
	head *Node
}

// Node for life
type Node struct {
	data     string
	nextNode *Node
}

func linkedList() {

}

func createList() *List {
	return &List{head: &Node{}}
}

func (l *List) tail(currentNode *Node) *Node {
	if currentNode.nextNode == nil {
		return currentNode
	}
	return l.tail(currentNode.nextNode)
}

func (l *List) insert(data string) {
	if l.head.data == "" {
		l.head.data = data
	} else {
		l.tail(l.head).nextNode = &Node{data: data}
	}
}
