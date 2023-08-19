package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstringCase1(t *testing.T) {
	expected := "bab"
	actual := longestPalindrome("babad")

	assert.Equal(t, expected, actual)
}

func TestLongestPalindromicSubstringCase2(t *testing.T) {
	expected := "bb"
	actual := longestPalindrome("cbbd")

	assert.Equal(t, expected, actual)
}

func TestLongestPalindromicSubstringCase3(t *testing.T) {
	expected := "a"
	actual := longestPalindrome("a")

	assert.Equal(t, expected, actual)
}
