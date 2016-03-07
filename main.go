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

func (l *List) insert(data string) {
	if l.head.data == "" {
		l.head.data = data
	}
}
