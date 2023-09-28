package piscine

func ListLast(l *List) interface{} {
	p := l.Head
	if p == nil {
		return nil
	}
	if p != nil {
		for p.Next != nil {
			p = p.Next
		}
	}
	return p.Data
}
