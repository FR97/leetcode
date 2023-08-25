package hard

func firstMissingPositive(nums []int) int {
	indexed := make([]bool, len(nums)+2)

	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v > 0 && v <= len(nums) {
			indexed[v] = true
		}
	}

	for i := 1; i < len(indexed); i++ {
		if !indexed[i] {
			return i
		}
	}

	return 0
}
