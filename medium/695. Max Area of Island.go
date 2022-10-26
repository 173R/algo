type Point struct {
	row int
	col int
}

func maxAreaOfIsland(grid [][]int) int {
	result := 0

	dimension := Point{
		row: len(grid),
		col: len(grid[0]),
	}

	//Точки которые уже были посещены при поиске частей острова
	piceOfIslands := map[Point]bool{} // R x C

	for row, _ := range grid { // R x C
		for col, _ := range grid[row] {
			//В каждой точке запускаем поиск в ширину
			maxArea := BFS(piceOfIslands, dimension, row, col, grid)
			if maxArea > result {
				result = maxArea
			}
		}
	}

	return result
}

func BFS(piceOfIslands map[Point]bool, dimension Point, row int, col int, grid [][]int) int {
	result := 0
	//Запускаем поиск только в случае если клетка подходит по условию и она не была ещё посещена(что заного не открывать остров несколько раз из разных клеток)
	if grid[row][col] == 1 && !piceOfIslands[Point{row, col}] {
		//Если подходит то увеличиваем размер острова который потом вернём и записываем клетку в посещённые
		result++
		piceOfIslands[Point{row, col}] = true

		//Дефолтный БФС по соседям
		if row-1 >= 0 {
			result += BFS(piceOfIslands, dimension, row-1, col, grid)
		}

		if row+1 < dimension.row {
			result += BFS(piceOfIslands, dimension, row+1, col, grid)
		}

		if col+1 < dimension.col {
			result += BFS(piceOfIslands, dimension, row, col+1, grid)
		}

		if col-1 >= 0 {
			result += BFS(piceOfIslands, dimension, row, col-1, grid)
		}

	}

	return result
}