package leetcode

func IsMatch(s string, p string) bool {
	m := make([][]bool, len(s)+1)
	for i := 0; i < len(m); i++ {
		m[i] = make([]bool, len(p)+1)
	}
	return isMatchDP(m,0, 0, s, p)
}

// Without DP
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	
	match := false
	if s != "" && (s[0] == p[0] || p[0] == '.') {
		match = true
	}
	
	if len(p) >= 2 && p[1] == '*' {
		return match && isMatch(s[1:], p) || isMatch(s, p[2:])
	} else {
		return match && isMatch(s[1:], p[1:])
	}
}

// With DP
func isMatchDP(m [][]bool, i, j int, s, p string) bool {
	if m[i][j] {
		return true
	}
	
	var ans bool
	
	if j == len(p) {
		ans = i == len(s)
	} else {
		match := false
		if i < len(s) && (s[i] == p[j] || p[j] == '.') {
			match = true
		}
		
		if j + 1 < len(p) && p[j+1] == '*' {
			ans = (match && isMatchDP(m, i+1, j, s, p)) || isMatchDP(m, i, j+2, s, p)
		} else {
			ans =  match && isMatchDP(m, i+1, j+1, s, p)
		}
	}
	
	m[i][j] = ans
	
	return ans
}
