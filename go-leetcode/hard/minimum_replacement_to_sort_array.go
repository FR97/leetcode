package hard

func minimumReplacement(nums []int) int64 {
	result := 0

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] > nums[i+1] {
			count := (nums[i] + nums[i+1] - 1) / nums[i+1]
			result += count - 1
			nums[i] = nums[i] / count
		}
	}

	return int64(result)
}
