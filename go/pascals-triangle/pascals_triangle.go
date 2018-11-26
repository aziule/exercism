package pascal

type triangle struct {
	size int
}

func Triangle(size int) [][]int {
	var rows [][]int

	for i := 0; i < size; i++ {
		if len(rows) == 0 {
			rows = append(rows, []int{1})
			continue
		}

		prev := rows[i-1]
		var row []int

		for j := 0; j <= i; j++ {
			var left int
			var right int

			if j == 0 {
				left = 0
				right = prev[j]
			} else if j == i {
				left = prev[j-1]
				right = 0
			} else {
				left = prev[j-1]
				right = prev[j]
			}

			row = append(row, left+right)
		}

		rows = append(rows, row)
	}

	return rows
}
