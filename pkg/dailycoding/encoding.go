package dailycoding

import (
	"strconv"
	"strings"
)

func NumDecodings(str string) int {
	digits := strings.Split(str, "")
	return waysOfEncodingDP(digits, make(map[string]int))
}

func waysOfEncodingDP(digits []string, dp map[string]int) int {
	if c, ok := dp[strings.Join(digits, "")]; ok {
		return c
	}

	if len(digits) >= 1 && digits[0] == "0" {
		return 0
	}

	if len(digits) == 0 {
		return 1
	}

	numberOfWays := 0
	numberOfWays = numberOfWays + waysOfEncodingDP(digits[1:], dp)

	if len(digits) == 2 {
		n, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			return -1
		}

		if n > 0 && n <= 26 {
			numberOfWays = numberOfWays + 1
		}
	}

	if len(digits) > 2 {
		numberOfWays = numberOfWays + waysOfEncodingDP(digits[2:], dp)
	}

	dp[strings.Join(digits, "")] = numberOfWays
	return numberOfWays
}

// Brute force solution
func waysOfEncoding(digits []string) int {
	if len(digits) == 0 {
		return 1
	}

	numberOfWays := 0
	numberOfWays = numberOfWays + waysOfEncoding(digits[1:])

	if len(digits) == 2 {
		n, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			return -1
		}

		if n > 0 && n <= 26 {
			numberOfWays = numberOfWays + 1
		}
	}

	if len(digits) > 2 {
		numberOfWays = numberOfWays + waysOfEncoding(digits[2:])
	}

	return numberOfWays
}

// Using string -
func waysDP(currentStr string, dp map[string]int) int {
	if c, ok := dp[currentStr]; ok {
		return c
	}

	if len(currentStr) == 0 {
		return 1
	}

	numberOfWays := 0
	numberOfWays = numberOfWays + waysDP(currentStr[1:], dp)

	if len(currentStr) == 2 {
		n, err := strconv.Atoi(currentStr)
		if err != nil {
			return -1
		}

		if n < 1 || n > 26 {
			return 0
		}

		numberOfWays = numberOfWays + 1
	}

	if len(currentStr) > 2 {
		numberOfWays = numberOfWays + waysDP(currentStr[2:], dp)
	}

	dp[currentStr] = numberOfWays
	return numberOfWays
}
