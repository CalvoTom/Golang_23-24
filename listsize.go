package piscine

func ListSize(l *List) int {
	var nb int
	var isTrue bool = true
	link := &List{}
	if l.Head == nil {
		return nb
	}
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
