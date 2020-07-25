package main

func numIslands(grid [][]byte) int {

	rows := len(grid)
	sum := 0

	if rows == 0 {
		return 0
	}

	columns := len(grid[0])

	if columns == 0 {
		return 0
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == '1' {
				sum++
				findNeighbours(grid, i, j)
			}
		}
	}

	return sum
}

func findNeighbours(grid [][]byte, i int, j int) {
	var x, y int

	// UP
	x = i - 1
	y = j
	if x >= 0 && grid[x][y] == '1' {
		grid[x][y] = '0'
		findNeighbours(grid, x, y)
	}

	// LEFT
	x = i
	y = j - 1
	if y >= 0 && grid[x][y] == '1' {
		grid[x][y] = '0'
		findNeighbours(grid, x, y)
	}

	// BOTTOM
	x = i + 1
	y = j
	if x < len(grid) && grid[x][y] == '1' {
		grid[x][y] = '0'
		findNeighbours(grid, x, y)
	}

	// RIGHT
	x = i
	y = j + 1
	if y < len(grid[0]) && grid[x][y] == '1' {
		grid[x][y] = '0'
		findNeighbours(grid, x, y)
	}
}
