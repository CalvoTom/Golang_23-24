package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validNode(root, -9223372036854775808, 9223372036854775807)
}

func validNode(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}
	if AtoiTest(node.Data) < min || AtoiTest(node.Data) > max {
		return false
	}
	return validNode(node.Left, min, AtoiTest(node.Data)-1) && validNode(node.Right, AtoiTest(node.Data)+1, max)
}

func AtoiTest(s string) int {
	if !IsValid(s) {
		return 0
	}
	sl := make([]int, 0)
	negative := false
	r := 0
	for _, v := range s {
		if IsValid(string(v)) && v >= 48 {
			sl = append(sl, int(v-48))
		} else if v == '-' {
			negative = true
		}
	}
	for i, v := range sl {
		r += v * Pow(10, len(sl)-i-1)
	}
	if negative {
		r = -r
	}
	return r
}

func IsValid(s string) bool {
	for i, v := range s {
		if v < rune('0') || v > rune('9') {
			if i == 0 && (s[i] == '+' || s[i] == '-') {
				continue
			}
			return false
		}
	}
	return true
}

func Pow(a int, b int) int {
	r := 1
	for i := 0; i < b; i++ {
		r *= a
	}
	return r
}
