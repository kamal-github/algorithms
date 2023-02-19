package leetcode

//Example 1:
//
//Input: ["flower","fow","flight"]
//Output: "fl"
//Example 2:
//
//Input: ["dog","racecar","car"]
//Output: ""
//Explanation: There is no common prefix among the input strings.

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	baseStr := strs[0]

	if len(strs) == 1 {
		return baseStr
	}

	var lcp string

	for i, ch := range baseStr {
		for _, str := range strs[1:] {
			if i >= len(str) || uint8(ch) != str[i] {
				return lcp
			}
		}

		lcp += string(ch)
	}

	return lcp
}

//func LongestCommonPrefix(strs []string) string {
//	cp := ""
//	mem := make(map[int32]int)
//	for i, c := range strs[0] {
//		mem[c] = i
//	}
//
//	for _, s := range strs[1:] {
//		tempPrefix := ""
//		tempPrefixL := 0
//		for i, c := range s {
//			if i == mem[c] {
//				tempPrefix = tempPrefix + string(c)
//				tempPrefixL++
//			}
//		}
//		if tempPrefixL <= len(cp) {
//			cp = tempPrefix
//		}
//	}
//	return cp
//}
