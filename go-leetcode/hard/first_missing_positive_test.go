package hard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstMissingPositiveCase1(t *testing.T) {
	expected := 3
	actual := firstMissingPositive([]int{1, 2, 0})
	assert.Equal(t, expected, actual)
}

func TestFirstMissingPositiveCase2(t *testing.T) {
	expected := 2
	actual := firstMissingPositive([]int{3, 4, -1, 1})
	assert.Equal(t, expected, actual)
}

func TestFirstMissingPositiveCase3(t *testing.T) {
	expected := 1
	actual := firstMissingPositive([]int{7, 8, 9, 11, 12})
	assert.Equal(t, expected, actual)
}
