package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	lp := l
	for p := 0; p < pos; p++ {
		if lp.Next == nil {
			return nil
		}
		lp = lp.Next
	}
	return lp
}
