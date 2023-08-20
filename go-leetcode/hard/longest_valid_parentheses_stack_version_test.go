package hard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestValidParenthesesStackCase1(t *testing.T) {
	expected := 2
	actual := longestValidParenthesesStack("(()")
	assert.Equal(t, expected, actual)
}

func TestLongestValidParenthesesStackCase2(t *testing.T) {
	expected := 4
	actual := longestValidParenthesesStack(")()())")
	assert.Equal(t, expected, actual)
}

func TestLongestValidParenthesesStackCase3(t *testing.T) {
	expected := 0
	actual := longestValidParenthesesStack("")
	assert.Equal(t, expected, actual)
}
