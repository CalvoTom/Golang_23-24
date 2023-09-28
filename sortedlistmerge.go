package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	current := n1
	if current == nil {
		n1 = n2
	} else {
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n2
	}
	ListSort(n1)
	return n1
}
