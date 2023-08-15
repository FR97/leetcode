package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesCase1(t *testing.T) {
	expected := 2
	actual := removeDuplicates([]int{1, 1, 2})
	assert.Equal(t, expected, actual)
}

func TestRemoveDuplicatesCase2(t *testing.T) {
	expected := 5
	actual := removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	assert.Equal(t, expected, actual)
}
