package pascal

func Triangle(size int) [][]int {
	if size == 1 {
		return [][]int{{1}}
	}

	prevRow := Triangle(size - 1)
	currRow := make([]int, size)

	currRow[0] = 1
	currRow[size-1] = 1

	for i := 1; i < size-1; i++ {
		currRow[i] = prevRow[size-2][i-1] + prevRow[size-2][i]
	}

	return append(prevRow, currRow)
}
