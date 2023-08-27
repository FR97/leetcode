package medium

func rotate(matrix [][]int) {
	lenI, lenJ := len(matrix), len(matrix[0])

	// transpose
	for i := 0; i < lenI; i++ {
		for j := i; j < lenJ; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ/2; j++ {
			matrix[i][j], matrix[i][lenJ-1-j] = matrix[i][lenJ-1-j], matrix[i][j]
		}
	}
}
