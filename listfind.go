package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	p := l.Head
	for !comp(ref, p.Data) {
		p = p.Next
	}
	return &p.Data
}
