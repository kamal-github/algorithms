package ctci

/*
Write a method to replace all spaces in a string with '%20'. You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string. (Note: If implementing in Java, please use a character array so that you can
perform this operation in place.)
EXAMPLE
Input: "Mr John Smith ", 13
Output: "Mr%20John%20Smith"
*/

// URLify is in place algorithm.
func URLify(s []byte, n int) []byte {
	// contains the index for last byte of `s`
	j := len(s) - 1

	// iterating from the original string(without whitespaces), moving each byte to
	// end of the `z`, if whitespace, replace it with `%20`.
	for i := n - 1; i >= 0; i-- {
		s[j] = s[i]
		if s[j] == ' ' {
			s[j] = '0'
			j--
			s[j] = '2'
			j--
			s[j] = '%'
		}
		j--
	}

	return s
}
