package dailycoding

func Swap(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b

	return a, b
}

func Swap2(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	return a, b
}
