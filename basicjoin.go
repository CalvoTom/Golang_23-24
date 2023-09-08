package piscine

func BasicJoin(elems []string) string {
	var string_join string
	for i := 0; i < len(elems); i++ {
		string_join += elems[i]
	}
	return string_join
}
