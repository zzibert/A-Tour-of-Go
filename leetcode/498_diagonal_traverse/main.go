package main

func main() {}

func findDiagonalOrder(matrix [][]int) []int {

	xLength := len(matrix)

	if xLength == 0 {
		return []int{}
	}

	yLength := len(matrix[0])

	return findDiagonalOrderHelper(matrix, []int{}, true, xLength, yLength, 0, 0)

}

func findDiagonalOrderHelper(matrix [][]int, array []int, goingUp bool, xLength, yLength, x, y int) []int {
	if isInside(x, y, xLength, yLength) {
		array = append(array, matrix[x][y])
	} else {
		return array
	}

	if goingUp {
		if isInside(x-1, y+1, xLength, yLength) {
			return findDiagonalOrderHelper(matrix, array, true, xLength, yLength, x-1, y+1)
		} else if isInside(x, y+1, xLength, yLength) {
			return findDiagonalOrderHelper(matrix, array, false, xLength, yLength, x, y+1)
		} else {
			return findDiagonalOrderHelper(matrix, array, false, xLength, yLength, x+1, y)
		}
	}

	if isInside(x+1, y-1, xLength, yLength) {
		return findDiagonalOrderHelper(matrix, array, false, xLength, yLength, x+1, y-1)
	} else if isInside(x+1, y, xLength, yLength) {
		return findDiagonalOrderHelper(matrix, array, true, xLength, yLength, x+1, y)
	}

	return findDiagonalOrderHelper(matrix, array, true, xLength, yLength, x, y+1)
}

func isInside(x, y, xLength, yLength int) bool {

	if x < 0 || y < 0 {
		return false
	}

	if x > xLength-1 || y > yLength-1 {
		return false
	}

	return true
}
