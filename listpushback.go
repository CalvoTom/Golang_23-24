package piscine

func ListPushBack(l *List, data interface{}) {
	NewNodel := NodeL{Data: data}
	if l.Head == nil {
		l.Head = &NewNodel
	} else {
		l.Tail.Next = &NewNodel
	}
	l.Tail = &NewNodel
}
