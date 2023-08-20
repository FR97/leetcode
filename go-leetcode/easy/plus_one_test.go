package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOneCase1(t *testing.T) {
	expected := []int{1, 2, 4}
	actual := plusOne([]int{1, 2, 3})
	assert.Equal(t, expected, actual)
}

func TestPlusOneCase2(t *testing.T) {
	expected := []int{4, 3, 2, 2}
	actual := plusOne([]int{4, 3, 2, 1})
	assert.Equal(t, expected, actual)
}

func TestPlusOneCase3(t *testing.T) {
	expected := []int{1}
	actual := plusOne([]int{0})
	assert.Equal(t, expected, actual)
}

func TestPlusOneCase5(t *testing.T) {
	expected := []int{1, 0}
	actual := plusOne([]int{9})
	assert.Equal(t, expected, actual)
}
