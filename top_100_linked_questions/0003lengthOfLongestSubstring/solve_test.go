package _003lengthOfLongestSubstring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
func Test_lengthOfLongestSubstring(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	})
	t.Run("example2", func(t *testing.T) {
		assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	})
	t.Run("example3", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	})
	t.Run("test case 138", func(t *testing.T) {
		assert.Equal(t, 2, lengthOfLongestSubstring("abba"))
	})
}

//Benchmark_solve-16    	 2106189	       569.4 ns/op	      48 B/op	       2 allocs/op
func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solve("abcabcbb")
	}
}

//Benchmark_solve2-16    	 9411247	       130.6 ns/op	       0 B/op	       0 allocs/op
func Benchmark_solve2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = solve2("abcabcbb")
	}
}
