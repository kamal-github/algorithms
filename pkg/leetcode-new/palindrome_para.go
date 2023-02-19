package main

import (
	"strings"
)

const (
	lowerCaseBeg = 97
	lowerCaseEnd = 122
	upperCaseBeg = 65
	upperCaseEnd = 90
	numBeg       = 48
	numEnd       = 57
)

func isPalindrome(s string) bool {
	b, e := 0, len(s)-1
	for b < e {
		if !isValidChar(s[b]) {
			b++
			continue
		}
		if !isValidChar(s[e]) {
			e--
			continue
		}
		if strings.ToLower(string(s[b])) != strings.ToLower(string(s[e])) {
			return false
		}
		b++
		e--
	}
	return true
}

func isValidChar(ch byte) bool {
	switch {
	case ch >= lowerCaseBeg && ch <= lowerCaseEnd:
	case ch >= upperCaseBeg && ch <= upperCaseEnd:
	case ch >= numBeg && ch <= numEnd:
	default:
		return false
	}

	return true
}
