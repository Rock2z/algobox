package _004findMedianSortedArrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		res := findMedianSortedArrays([]int{1, 3}, []int{2})
		assert.EqualValues(t, 2, res)
	})
	t.Run("example1", func(t *testing.T) {
		res := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
		assert.EqualValues(t, 2.5, res)
	})
	t.Run("test case 2", func(t *testing.T) {
		res := findMedianSortedArrays([]int{0, 0}, []int{0, 0})
		assert.EqualValues(t, 0, res)
	})
}
