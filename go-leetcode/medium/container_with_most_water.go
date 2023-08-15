package medium

func maxArea(height []int) int {
	maxLeft := 0
	maxRight := len(height) - 1
	area := 0

	for maxLeft < maxRight {
		area = max(area, (maxRight-maxLeft)*min(height[maxLeft], height[maxRight]))
		if height[maxLeft] < height[maxRight] {
			maxLeft++
		} else {
			maxRight--
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if b > a {
		return a
	}
	return b
}
