type Point struct {
	x int
	y int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	queue := []Point{{sc, sr}}
	rows := len(image)
	cols := len(image[0])

	baseColor := image[sr][sc]

	for len(queue) > 0 {
		currentPoint := slicePopFirst(&queue)

		if image[currentPoint.y][currentPoint.x] != color {
			image[currentPoint.y][currentPoint.x] = color

			top := Point{currentPoint.x, currentPoint.y - 1}
			bottom := Point{currentPoint.x, currentPoint.y + 1}
			left := Point{currentPoint.x - 1, currentPoint.y}
			right := Point{currentPoint.x + 1, currentPoint.y}

			if top.y >= 0 && image[top.y][top.x] == baseColor {
				queue = append(queue, top)
			}

			if bottom.y < rows && image[bottom.y][bottom.x] == baseColor {
				queue = append(queue, bottom)
			}

			if left.x >= 0 && image[left.y][left.x] == baseColor {
				queue = append(queue, left)
			}

			if right.x < cols && image[right.y][right.x] == baseColor {
				queue = append(queue, right)
			}

		}
	}

	return image
}

func slicePopFirst(queuePtr *[]Point) Point {
	result := (*queuePtr)[0]
	*queuePtr = (*queuePtr)[1:]
	return result

}