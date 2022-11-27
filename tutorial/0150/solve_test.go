package _150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	assert.Equal(t, 6, evalRPN([]string{"4", "13", "5", "/", "+"}))
}
