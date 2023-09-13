package easy

import (
	"github.com/FR97/leetcode/go-leetcode/medium"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonPrefixCase1(t *testing.T) {
	expected := "fl"
	actual := medium.longestCommonPrefix([]string{"flower", "flow", "flight"})
	assert.Equal(t, expected, actual)
}

func TestLongestCommonPrefixCase2(t *testing.T) {
	expected := ""
	actual := medium.longestCommonPrefix([]string{"dog", "racecar", "car"})
	assert.Equal(t, expected, actual)
}
