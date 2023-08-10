package easy

func twoSum(nums []int, target int) []int {
	store := map[int]int{}
	for i, v := range nums {
		key := target - v
		if j, ok := store[key]; ok {
			return []int{j, i}
		}
		store[v] = i
	}
	return []int{}
}
