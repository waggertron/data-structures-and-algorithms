package linkedlists

type Node struct {
	Value any
	Next  *Node
}

type Linkedlist struct {
	Head   *Node
	Length int
}

func (l *Linkedlist) Add(value any) {
	n := &Node{Value: value}

	cur := l.Head
	if cur == nil {
		l.Head = n
		l.Length = 1
		return
	}

	for cur != nil {
		next := cur.Next
		if next != nil {
			cur = next
		}
	}

	cur.Next = n

	l.Length += 1
}

func (l *Linkedlist) Remove(value any) {
	cur := l.Head
	if cur == nil {
		return
	}
	next := cur.Next
	for {
		if next == nil {
			return
		}
		if next.Value == value {
			break
		}
		cur = cur.Next
		next = cur.Next
	}

	cur.Next = next.Next

	l.Length -= 1
}
