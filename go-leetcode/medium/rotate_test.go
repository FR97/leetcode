package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateCase1(t *testing.T) {
	expected := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	actual := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(actual)

	assert.Equal(t, expected, actual)
}

func TestRotateCase2(t *testing.T) {
	expected := [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}

	actual := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(actual)

	assert.Equal(t, expected, actual)
}
