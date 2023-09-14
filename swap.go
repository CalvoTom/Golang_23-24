package piscine

func Swap(a *int, b *int) {
	val_a := *a
	val_b := *b
	*a = val_b
	*b = val_a
}
