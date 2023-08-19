package medium

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 0
	for i := 0; i < len(s); i++ {

		startOdd, maxLenOdd := spread(s, i, i)
		startEven, maxLenEven := spread(s, i, i+1)

		if maxLenOdd > maxLenEven && maxLenOdd > maxLen {
			maxLen = maxLenOdd
			start = startOdd
		} else if maxLenEven > maxLen {
			maxLen = maxLenEven
			start = startEven
		}

	}

	return s[start : start+maxLen+1]
}

func spread(s string, l int, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return l + 1, (r - 1) - (l + 1)
}
