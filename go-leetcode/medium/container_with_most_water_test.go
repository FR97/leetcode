package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxAreaCase1(t *testing.T) {
	expected := 49
	actual := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(t, expected, actual)
}

func TestMaxAreaCase2(t *testing.T) {
	expected := 1
	actual := maxArea([]int{1, 1})
	assert.Equal(t, expected, actual)
}
