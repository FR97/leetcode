package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsertCase1(t *testing.T) {
	expected := 2
	actual := searchInsert([]int{1, 3, 5, 6}, 5)
	assert.Equal(t, expected, actual)
}

func TestSearchInsertCase2(t *testing.T) {
	expected := 1
	actual := searchInsert([]int{1, 3, 5, 6}, 2)
	assert.Equal(t, expected, actual)
}

func TestSearchInsertCase3(t *testing.T) {
	expected := 4
	actual := searchInsert([]int{1, 3, 5, 6}, 7)
	assert.Equal(t, expected, actual)
}
