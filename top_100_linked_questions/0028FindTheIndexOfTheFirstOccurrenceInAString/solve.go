package _028FindTheIndexOfTheFirstOccurrenceInAString

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	next := findNext(needle)

	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		op := haystack[i : i+1]
		tp := needle[j : j+1]
		if op != tp {
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}
		} else if j+1 == len(needle) {
			return i - j
		} else {
			i++
			j++
		}
	}
	return -1
}

func findNext(s string) []int {
	l := len(s)
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		ks := s[:i+1]
		fmt.Println("ks=", ks)
		for j := 1; j <= i; j++ {
			prefix := ks[:j]
			suffix := ks[len(ks)-j:]
			fmt.Println("prefix=", prefix, "suffix=", suffix)
			if prefix == suffix {
				ret[i] = j
			}
		}
	}
	return ret
}

//         a d d a d d
//next  -1 0 1 3 0 1 0
//
//		 a a b a a f
//next   0 1 0 1 2 0
//
//		a a b a a b a a f a
//		a a b a a f
//			  a a b a a f a

//i=0,j=6
//i=3,j=3

//mississippi
// issip
//	issip
