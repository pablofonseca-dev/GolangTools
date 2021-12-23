package utils

func SwapValues(a *float64, b *float64) {
	var temp float64 = *a
	*a = *b
	*b = temp
}
