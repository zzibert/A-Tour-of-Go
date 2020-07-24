package main

type Coordinate struct {
	x int
	y int
}

const INF = 2147483647

func wallsAndGates(rooms [][]int) {

	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	rows := len(rooms)
	columns := len(rooms[0])

	queue := make([]Coordinate, 0)

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			cell := rooms[i][j]
			if cell == 0 {
				queue = append(queue, Coordinate{i, j})
			}
		}
	}

	for len(queue) > 0 {
		zero := queue[0]
		queue = queue[1:]
		incrementNeighbours(rooms, zero, rows, columns, &queue)
	}
}

func incrementNeighbours(rooms [][]int, zero Coordinate, rows int, columns int, queue *[]Coordinate) {
	var x, y int

	// UP
	x = zero.x - 1
	y = zero.y
	if x >= 0 && rooms[x][y] > rooms[zero.x][zero.y] {
		rooms[x][y] = rooms[zero.x][zero.y] + 1
		*queue = append(*queue, Coordinate{x, y})
	}
	// LEFT
	x = zero.x
	y = zero.y - 1
	if y >= 0 && rooms[x][y] > rooms[zero.x][zero.y] {
		rooms[x][y] = rooms[zero.x][zero.y] + 1
		*queue = append(*queue, Coordinate{x, y})
	}
	// DOWN
	x = zero.x + 1
	y = zero.y
	if x < rows && rooms[x][y] > rooms[zero.x][zero.y] {
		rooms[x][y] = rooms[zero.x][zero.y] + 1
		*queue = append(*queue, Coordinate{x, y})
	}
	// RIGHT
	x = zero.x
	y = zero.y + 1
	if y < columns && rooms[x][y] > rooms[zero.x][zero.y] {
		rooms[x][y] = rooms[zero.x][zero.y] + 1
		*queue = append(*queue, Coordinate{x, y})
	}
}
