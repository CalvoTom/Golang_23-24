package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := NodeI{Data: data_ref}
	if l == nil {
		return &newNode
	} else if l.Data > data_ref {
		newNode.Next = l
		return &newNode
	}
	for current := l; current != nil && current.Data < data_ref; current = current.Next {
		if current.Next != nil && current.Next.Data >= data_ref {
			newNode.Next = current.Next
			current.Next = &newNode
		} else if current.Next == nil && current.Data < data_ref {
			current.Next = &newNode
		}
	}
	return l
}
