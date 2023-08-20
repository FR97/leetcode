package hard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestValidParenthesesCase1(t *testing.T) {
	expected := 2
	actual := longestValidParentheses("(()")
	assert.Equal(t, expected, actual)
}

func TestLongestValidParenthesesCase2(t *testing.T) {
	expected := 4
	actual := longestValidParentheses(")()())")
	assert.Equal(t, expected, actual)
}

func TestLongestValidParenthesesCase3(t *testing.T) {
	expected := 0
	actual := longestValidParentheses("")
	assert.Equal(t, expected, actual)
}
