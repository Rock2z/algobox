package _242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
}
