package _977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
	})
}
