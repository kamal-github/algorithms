package leetcode

func LongestSubstringLength(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	var longest, currentLongest, next int
	occurrenceCache := make(map[int32]int)
	
	for i, ch := range s {
		pos, found := occurrenceCache[ch]
		// m is tracking the m character to restart the find the longest substring
		// and the positions before the m index is irrelevant, as it is already been tracked.
		if found && pos >= next {
			// update the currentLongest in order to avoid the recalculation(optimization step)
			currentLongest = i - pos
			// update m to the m character of previously found ch
			next = occurrenceCache[ch] + 1
		} else {
			// when characters are not found in cache
			currentLongest++
		}
		// updating the occurrence index
		occurrenceCache[ch] = i
		
		// Keep updating the final longest number
		longest = max(currentLongest, longest)
	}
	
	return longest
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
