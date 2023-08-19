package medium

import "math"

func reverse(x int) int {
	reversed := 0

	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10

		if reversed > math.MaxInt32/10 || reversed < math.MinInt32/10 {
			return 0
		}
	}

	return reversed
}
