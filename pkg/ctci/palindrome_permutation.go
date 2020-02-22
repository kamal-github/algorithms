package ctci

func PalindromePermutation(s []byte) bool {
	cMap := make(map[byte]int)

	for _, b := range s {
		if b == ' ' {
			continue
		}

		if c, ok := cMap[b]; ok {
			cMap[b] = c + 1
		} else {
			cMap[b] = 1
		}
	}

	oddCount := 0

	for _, v := range cMap {
		if v%2 == 1 {
			oddCount++
		}
	}
	if oddCount > 1 {
		return false
	}

	return true
}
