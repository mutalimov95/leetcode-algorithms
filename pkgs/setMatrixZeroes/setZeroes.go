package setMatrixZeroes

func setZeroes(matrix [][]int) {
	rowMap := make(map[int]struct{})
	colMap := make(map[int]struct{})
	for i, row := range matrix {
		for j, cell := range row {
			if cell == 0 {
				rowMap[i] = struct{}{}
				colMap[j] = struct{}{}
			}
		}
	}
	for r := range rowMap {
		zereableRow(matrix, r)
	}
	for c := range colMap {
		zereableColumn(matrix, c)
	}
}

func zereableColumn(matrix [][]int, number int) {
	for _, row := range matrix {
		row[number] = 0
	}
}

func zereableRow(matrix [][]int, number int) {
	for j := range matrix[number] {
		matrix[number][j] = 0
	}
}
