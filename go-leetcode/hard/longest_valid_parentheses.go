package hard

func longestValidParentheses(s string) int {
	mem := make([]int, len(s))
	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				mem[i] = 2
				if i-2 >= 0 {
					mem[i] += mem[i-2]
				}
			} else if i-mem[i-1]-1 >= 0 && s[i-mem[i-1]-1] == '(' {
				mem[i] = mem[i-1] + 2
				if i-mem[i-1]-2 >= 0 {
					mem[i] += mem[i-mem[i-1]-2]
				}
			}
			if mem[i] > max {
				max = mem[i]
			}
		}
	}
	return max
}
