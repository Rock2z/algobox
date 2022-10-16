package _028FindTheIndexOfTheFirstOccurrenceInAString

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findNext(t *testing.T) {
	t.Run("self test", func(t *testing.T) {
		s := "aabaaf"
		fmt.Println(findNext(s))
	})
}

func Test_strStr(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		assert.Equal(t, 0, strStr("sadbutsad", "sad"))
	})
	t.Run("demo2", func(t *testing.T) {
		assert.Equal(t, -1, strStr("leetcode", "leeto"))
	})
	t.Run("demo_isaac", func(t *testing.T) {
		assert.Equal(t, 3, strStr("aabaabaafa", "aabaaf"))
	})
	t.Run("demo9", func(t *testing.T) {
		assert.Equal(t, -1, strStr("aaa", "aaaa"))
	})
	t.Run("demo15", func(t *testing.T) {
		assert.Equal(t, -1, strStr("mississippi", "issipi"))
	})
}
