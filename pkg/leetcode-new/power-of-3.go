package main

func isPowerOfThree(n int) bool {
	if n == 0 || n == -1 {
		return false
	}
	for n != 1 {
		n = n / 3
		if (n > 1 && n < 3) || (n < -1 && n > -3) {
			return false
		}
	}

	return true
}
