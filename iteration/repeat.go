package iteration

func Repeat(symbol string, count int) string {
	var finalSymbol string
	for i := 0; i < count; i++ {
		finalSymbol += symbol
	}
	return finalSymbol
}
