package medium

func permute(nums []int) [][]int {
	var result [][]int

	if len(nums) == 1 {
		result = append(result, nums)
		return result
	}

	for i := range nums {
		toPermute := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		for _, permuted := range permute(toPermute) {
			result = append(result, append([]int{nums[i]}, permuted...))
		}
	}

	return result
}
