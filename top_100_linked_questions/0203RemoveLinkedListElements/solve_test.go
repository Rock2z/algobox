package _203RemoveLinkedListElements

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"algobox/top_100_linked_questions/common"
)

func Test_removeElements(t *testing.T) {
	t.Run("demo_1", func(t *testing.T) {
		ret := removeElements(common.MakeNodeFromSlice([]int{1, 2, 6, 3, 4, 5, 6}), 6)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, common.MakeSliceFromNode(ret))
	})
	t.Run("demo_2", func(t *testing.T) {
		ret := removeElements(common.MakeNodeFromSlice([]int{}), 1)
		assert.Len(t, common.MakeSliceFromNode(ret), 0)
	})
	t.Run("demo_3", func(t *testing.T) {
		ret := removeElements(common.MakeNodeFromSlice([]int{7, 7, 7, 7}), 7)
		assert.Len(t, common.MakeSliceFromNode(ret), 0)
	})
	t.Run("demo_50", func(t *testing.T) {
		ret := removeElements(common.MakeNodeFromSlice([]int{1, 2, 2, 1}), 2)
		assert.Equal(t, []int{1, 1}, common.MakeSliceFromNode(ret))
	})
}
