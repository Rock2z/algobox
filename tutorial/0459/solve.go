package _459

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	b := []byte(s)

	for i := 1; i <= len(b)/2; i++ {
		p := b[0:i]
		for j := i; j < len(b); j += i {
			if j+i > len(b) || string(b[j:j+i]) != string(p) {
				break
			}
			if j+i == len(b) {
				return true
			}
		}
	}
	return false
}

func repeatedSubstringPattern2(s string) bool {
	return strings.Contains((s + s)[1:len(s)*2-1], s)
}
