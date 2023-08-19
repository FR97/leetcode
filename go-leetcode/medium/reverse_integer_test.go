package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseIntegerCase1(t *testing.T) {
	expected := 321
	actual := reverse(123)

	assert.Equal(t, expected, actual)
}

func TestReverseIntegerCase2(t *testing.T) {
	expected := -321
	actual := reverse(-123)

	assert.Equal(t, expected, actual)
}

func TestReverseIntegerCase3(t *testing.T) {
	expected := 21
	actual := reverse(120)

	assert.Equal(t, expected, actual)
}
