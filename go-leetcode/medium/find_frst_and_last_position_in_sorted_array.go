package medium

func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		j := len(nums) - i - 1
		if nums[i] == target && start == -1 {
			start = i
		}
		if nums[j] == target && end == -1 {
			end = j
		}
		if (start != -1 && end != -1) || (start != -1 && nums[j] < target) || (end != -1 && nums[i] > target) {
			return []int{start, end}
		}
	}
	return []int{start, end}
}
