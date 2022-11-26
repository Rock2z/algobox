package _015

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum2(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}, threeSum2([]int{-1, 0, 1, 2, -1, -4}))
	})
	t.Run("demo2", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{-2, 0, 2},
		}, threeSum2([]int{-2, 0, 0, 2, 2}))
	})
}
