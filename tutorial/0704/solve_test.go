package _704

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	})
	t.Run("demo2", func(t *testing.T) {
		assert.Equal(t, -1, search([]int{2 - 1, 0, 3, 5, 9, 12}, 2))
	})
}
