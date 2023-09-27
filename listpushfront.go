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
	switch {
	case l.Head == nil:
		l.Head = &NewNodel
	default:
		NewNodel.Next = l.Head
		l.Head = &NewNodel
	}
}
