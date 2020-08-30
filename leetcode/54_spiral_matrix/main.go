package main

func main() {}

func spiralOrder(matrix [][]int) []int {

	xLength := len(matrix)

	if xLength == 0 {
		return []int{}
	}

	yLength := len(matrix[0])

	checkedMatrix := make([][]int, xLength)

	for i := 0; i < xLength; i++ {
		checkedMatrix[i] = make([]int, yLength)
	}

	var x, y, counter int
	array := make([]int, 0)

	xAxis := 0
	yAxis := 1

	for {
		array = append(array, matrix[x][y])
		checkedMatrix[x][y] = 1
		for {
			if counter > 4 {
				return array
			}
			if isInside(x+xAxis, y+yAxis, xLength, yLength) && checkedMatrix[x+xAxis][y+yAxis] == 0 {
				counter = 0
				break
			}
			xAxis, yAxis = changeCourse(xAxis, yAxis)
			counter++
		}
		x += xAxis
		y += yAxis
	}

	return array

}

func changeCourse(xAxis, yAxis int) (int, int) {
	if xAxis == 0 && yAxis == 1 {
		return 1, 0
	}

	if xAxis == 1 && yAxis == 0 {
		return 0, -1
	}

	if xAxis == 0 && yAxis == -1 {
		return -1, 0
	}

	return 0, 1
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
