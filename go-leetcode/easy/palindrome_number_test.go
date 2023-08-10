package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeCase1(t *testing.T) {
	assert.True(t, isPalindrome(121), "number should be palindromme")
}

func TestIsPalindromeCase2(t *testing.T) {
	assert.False(t, isPalindrome(-121), "number should not be palindrome")
}

func TestIsPalindromeCase3(t *testing.T) {
	assert.False(t, isPalindrome(10), "number should not be palindrome")
}
