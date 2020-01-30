package leetcode

// O(n^3)
func LongestPalindrome(s string) string {
	inpLength := len(s)
	if inpLength <= 1 {
		return s
	}
	largest := s[0:1]
	maxSize := 1
	for i:=0; i<inpLength-1; i++ {
		for j:=i+1;j<inpLength;j++{
			if s[j] == s[i] && isPalindrome(s, i, j) {
				if j-i+1 > maxSize {
					maxSize = j-i+1
					largest = s[i:j+1]
				}
			}
		}
	}

	return largest
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// O(n^2)
func LongestPalindromeDP(s string) string {
	inpLength := len(s)
	if inpLength <= 1 {
		return s
	}

	mem := make([][]bool, inpLength)
	for i:=0; i<inpLength; i++{
		mem[i] = make([]bool, inpLength)
		mem[i][i] = true
	}

	largest := s[0:1]
	for i:=0; i<=inpLength-2; i++{
		if s[i] == s[i+1] {
			mem[i][i+1] = true
			largest = s[i:i+2]
		}
	}

	for i:=3; i <= inpLength; i++ {
		for j:=0; j<=inpLength-i; j++ {
			k := i+j-1
			if mem[j+1][k-1] && s[j] == s[k]{
				mem[j][k] = true
				largest = s[j:k+1]
			}
		}
	}

	return largest
}

