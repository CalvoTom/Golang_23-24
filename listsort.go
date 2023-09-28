package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	} else {
		for current := l; current != nil; current = current.Next {
			index := current.Next
			for index != nil {
				if current.Data > index.Data {
					temp := current.Data
					current.Data = index.Data
					index.Data = temp
				}
				index = index.Next
			}
		}
	}
	return l
}
