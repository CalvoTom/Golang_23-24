package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	NewNodel := NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = &NewNodel
	} else {
		l.Head.Next = l.Head
	}
	l.Head = &NewNodel
}
