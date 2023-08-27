package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermutationsCase1(t *testing.T) {
	expected := [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
		{2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	actual := permute([]int{1, 2, 3})

	assert.Equal(t, expected, actual)
}

func TestPermutationsCase2(t *testing.T) {
	expected := [][]int{{0, 1}, {1, 0}}
	actual := permute([]int{0, 1})

	assert.Equal(t, expected, actual)
}

func TestPermutationsCase3(t *testing.T) {
	expected := [][]int{{1}}
	actual := permute([]int{1})

	assert.Equal(t, expected, actual)
}
