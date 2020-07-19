package main

func searchMatrix(matrix [][]int, target int) bool {

	rows := len(matrix)

	if rows == 0 {
		return false
	}

	columns := len(matrix[0])

	if columns == 0 {
		return false
	}

	var halfRows, halfColumns int

	if rows%2 == 0 {
		halfRows = (rows / 2) - 1
	} else {
		halfRows = rows / 2
	}

	if columns%2 == 0 {
		halfColumns = (columns / 2) - 1
	} else {
		halfColumns = columns / 2
	}

	if rows == 1 && columns == 1 {
		return matrix[0][0] == target
	}

	if rows <= 3 && columns <= 3 {
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if matrix[i][j] == target {
					return true
				}
			}
		}
	}

	pivot := matrix[halfRows][halfColumns]
	if pivot == target {
		return true
	}
	if pivot < target {
		return searchMatrix(matrix[:halfRows][(halfColumns):], target) ||
			searchMatrix(matrix[(halfRows):][halfColumns:], target) ||
			searchMatrix(matrix[halfRows:][:halfColumns], target)
	}

	return searchMatrix(matrix[:halfRows][(halfColumns):], target) ||
		searchMatrix(matrix[halfRows:][:halfColumns], target) ||
		searchMatrix(matrix[halfRows:][:halfColumns], target)

}
