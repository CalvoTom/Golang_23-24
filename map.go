package piscine

func Map(f func(int) bool, a []int) []bool {
	var arrayReturned []bool
	for _, val := range a {
		arrayReturned = append(arrayReturned, f(val))
	}
	return arrayReturned
}
