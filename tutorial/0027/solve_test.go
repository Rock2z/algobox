package _027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		assert.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 3))
	})
	t.Run("demo2", func(t *testing.T) {
		assert.Equal(t, 5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	})
}
