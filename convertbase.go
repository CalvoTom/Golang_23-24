package piscine

func atoiBase(s string, base string) int {
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

func printNbrBase(nbr int, base string) string {
	var m int
	var bn int = len(base)
	sn := ""
	fn := ""
	if nbr >= 0 {
		m = nbr
	} else {
		m = nbr * (-1)
	}
	for e := 0; e == 0; e = e + 0 {
		if m/bn >= 1 {
			sn += string(base[m%bn])
			m = m / bn
		} else {
			sn += string(base[m%bn])
			e = e + 1
		}
	}
	for d := len(sn) - 1; d >= 0; d-- {
		fn += string(sn[d])
	}
	return fn
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := atoiBase(nbr, baseFrom)
	return printNbrBase(n, baseTo)
}
