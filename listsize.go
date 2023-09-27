package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	var nb int
	var isTrue bool = true
	link := &List{}
	for isTrue {
		nb += 1
		ListPushFront(link, l.Head)
		l.Head = l.Head.Next
		if l.Head == nil {
			isTrue = false
		}
	}
	return nb
}
