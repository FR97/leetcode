package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefixCase1(t *testing.T) {
	expected := "fl"
	actual := longestCommonPrefix([]string{"flower", "flow", "flight"})
	assert.Equal(t, expected, actual)
}

func TestLongestCommonPrefixCase2(t *testing.T) {
	expected := ""
	actual := longestCommonPrefix([]string{"dog", "racecar", "car"})
	assert.Equal(t, expected, actual)
}
