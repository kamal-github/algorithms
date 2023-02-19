package main

// "abcbbcbb"
func lengthOfLongestSubstring(s string) int {
	maxSubStrLen := 0
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	charsFreq := make(map[uint8]int)

	i := 0
	j := 0

	localLongestSubStrLen := 0
	for i < len(s) {
		for j < len(s) {
			if v, ok := charsFreq[s[j]]; !ok || v == 0 {
				charsFreq[s[j]]++
				j++
				localLongestSubStrLen++
				continue
			}
			break
		}
		if maxSubStrLen < localLongestSubStrLen {
			maxSubStrLen = localLongestSubStrLen
		}

		// preparing for next iteration
		charsFreq[s[i]]--
		i++
		localLongestSubStrLen = 0
	}

	return maxSubStrLen
}
