package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(node *TreeNode, data string) *TreeNode {
	if node == nil {
		return &TreeNode{Data: data}
	}
	if data < node.Data {
		node.Left = BTreeInsertData(node.Left, data)
		node.Left.Parent = node
	} else {
		node.Right = BTreeInsertData(node.Right, data)
		node.Right.Parent = node
	}
	return node
}
