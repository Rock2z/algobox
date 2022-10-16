package _002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commonChars(t *testing.T) {
	assert.ElementsMatch(t, []string{"e", "l", "l"}, commonChars([]string{"bella", "label", "roller"}))
	assert.ElementsMatch(t, []string{"c", "o"}, commonChars([]string{"cool", "lock", "cook"}))
}
