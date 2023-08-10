package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumCase1(t *testing.T) {
	expected := []int{0, 1}
	actual := twoSum([]int{2, 7, 11, 15}, 9)

	assert.Equal(t, expected, actual, "values don't match")
}

func TestTwoSumCase2(t *testing.T) {
	expected := []int{1, 2}
	actual := twoSum([]int{3, 2, 4}, 6)

	assert.Equal(t, expected, actual, "values don't match")
}
