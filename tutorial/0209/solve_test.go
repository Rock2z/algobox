package _209

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minSubArrayLen(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	})
	t.Run("demo2", func(t *testing.T) {
		assert.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))
	})
	t.Run("demo3", func(t *testing.T) {
		assert.Equal(t, 0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	})
	t.Run("demo4", func(t *testing.T) {
	})
}
