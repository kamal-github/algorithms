package ctci

import "math"

/*
There are three types of edits that can be performed on strings: insert a character,
remove a character, or replace a character. Given two strings, write a function to check if they are
one edit (or zero edits) away.
EXAMPLE
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false
*/
func OneEditAway(s1, s2 []byte) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1.0 {
		return false
	}

	iMap := make(map[byte]int)

	for i, b := range s1 {
		iMap[b] = i
	}

	diff := 0

	for i, b := range s2 {
		j, ok := iMap[b]
		if !ok {
			diff++
		} else if ok && math.Abs(float64(i-j)) > 1.0 {
			diff++
		}
	}

	if diff > 1 {
		return false
	}
	return true
}

// a-b-c-a-b-d
