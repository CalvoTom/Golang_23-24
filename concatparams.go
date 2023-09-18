package piscine

func ConcatParams(args []string) string {
	var str string
	for _, val := range args {
		str += val + "\n"
	}
	return str
}
