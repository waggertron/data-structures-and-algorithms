package linkedlists

type Node struct {
	Value any
	Next  *Node
}

type Linkedlist struct {
	Head   *Node
	Length int
}

func (l *Linkedlist) Add(n Node) {
	cur := l.Head
	if cur == nil {
		l.Head = &n
		return
	}

	for cur != nil {
		next := cur.Next
		if next != nil {
			cur = next
		}
	}

	cur.Next = &n
}

