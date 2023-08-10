package easy

func isPalindrome(x int) bool {
	num := x
	remainder := 0
	val := 0
	for num > 0 {
		remainder = num % 10
		val = val*10 + remainder
		num = num / 10
	}

	if x == val {
		return true
	}

	return false
}
