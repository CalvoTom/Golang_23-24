package piscine

func Sqrt(nb int) int {
	var rac int
	for i := 1; i < nb; i++ {
		rac = nb / i
		if rac*rac == nb {
			return rac
		}
	}
	return 0
}
