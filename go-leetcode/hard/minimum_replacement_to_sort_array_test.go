package hard

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumReplacementCase1(t *testing.T) {
	expected := int64(0)
	actual := minimumReplacement([]int{1, 2, 3, 4, 5})
	assert.Equal(t, expected, actual)
}
