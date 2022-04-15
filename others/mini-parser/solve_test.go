package mini_parser

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

//https://leetcode-cn.com/problems/mini-parser/
func Test_parse(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		assert.EqualValues(
			t, "324",
			displayResult(deserialize("324")),
		)
	})
	t.Run("example2", func(t *testing.T) {
		assert.EqualValues(
			t, "[123,[456,[789]]]",
			displayResult(deserialize("[123,[456,[789]]]")),
		)
	})
	t.Run("self test - negative num", func(t *testing.T) {
		assert.EqualValues(
			t, "[123,[-456,[789]]]",
			displayResult(deserialize("[123,[-456,[789]]]")),
		)
	})
	t.Run("test case 2", func(t *testing.T) {
		assert.EqualValues(
			t, "[123,456,[788,799,833],[[]],10,[]]",
			displayResult(deserialize("[123,456,[788,799,833],[[]],10,[]]")),
		)
	})
	t.Run("test case 28", func(t *testing.T) {
		res := deserialize("[0]")
		assert.EqualValues(
			t, "[0]",
			displayResult(res),
		)
	})
}

func displayResult(nest *NestedInteger) string {
	if nest.IsInteger() {
		return strconv.Itoa(nest.GetInteger())
	}
	p := "["
	for i, n := range nest.GetList() {
		p += displayResult(n)
		if i < len(nest.GetList())-1 {
			p += ","
		}
	}
	p += "]"
	return p
}
