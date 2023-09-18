package piscine

func ConcatParams(args []string) string {
	var str string
	for i, val := range args {
		if i != len(args)-1 {
			str += val + "\n"
		} else {
			str += val
		}
	}
	return str
}
