package _459RepeatedSubstringPattern

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		sub := s[0 : i+1]
		if len(s)%len(sub) != 0 {
			continue
		}
		if strings.Repeat(sub, len(s)/len(sub)) == s {
			return true
		}
	}
	return false
}
