package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchRangeCase1(t *testing.T) {
	expected := []int{3, 4}
	actual := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	assert.Equal(t, expected, actual)
}

func TestSearchRangeCase2(t *testing.T) {
	expected := []int{-1, -1}
	actual := searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	assert.Equal(t, expected, actual)
}

func TestSearchRangeCase3(t *testing.T) {
	expected := []int{-1, -1}
	actual := searchRange([]int{}, 0)
	assert.Equal(t, expected, actual)
}
