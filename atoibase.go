package piscine

func AtoiBase(s string, base string) int {
	var m int
	var bn int = len(base)
	puis := 1
	comp := 0

	if len(base) <= 1 {
		return 0
	} else {
		for i := 0; i < len(base); i++ {
			if base[i] == 43 || base[i] == 45 {
				return 0
			} else {
				for j := i + 1; j < len(base); j++ {
					if base[i] == base[j] {
						return 0
					}
				}
			}
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for c := 0; c < bn; c++ {
			if s[i] == base[c] {
				if comp == 0 {
					m += c
					comp = comp + 1
				} else {
					puis = puis * bn
					m += c * puis
				}
			}
		}
	}
	return m
}
