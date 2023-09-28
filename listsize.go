package piscine

func ListSize(l *List) int {
	count := 0
	p := l.Head
	if p != nil {
		count = 1
		for p.Next != nil {
			count++
			p = p.Next
		}
	}
	return count
}
