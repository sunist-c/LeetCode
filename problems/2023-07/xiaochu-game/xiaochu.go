package xiaochu_game

var (
	directions = []point{
		directionsUp, directionsDown, directionsLeft, directionsRight,
	}

	directionsUp = point{
		x: 1,
		y: 0,
	}
	directionsDown = point{
		x: -1,
		y: 0,
	}
	directionsLeft = point{
		x: 0,
		y: -1,
	}
	directionsRight = point{
		x: 0,
		y: 1,
	}
)

type point struct {
	x int
	y int
}

type Chessboard struct {
	board [][]int
	has   map[point]struct{}
}

func (c *Chessboard) transform(p point) (result point) {
	return point{
		x: len(c.board) - p.x - 1,
		y: p.y,
	}
}

func (c *Chessboard) findAll(minLength int) (needContinue bool) {
	needContinue = false
	for p := range c.has {
		indexes := c.find(p)
		if c.remove(minLength, indexes) {
			needContinue = true
		}
	}

	return needContinue
}

func (c *Chessboard) find(p point) (indexes []point) {
	queue, value := []point{p}, c.board[p.x][p.y]

	if value == 0 {
		return []point{}
	}

	var visited = make([][]bool, len(c.board))
	for i := 0; i < len(c.board); i++ {
		visited[i] = make([]bool, len(c.board[i]))
	}

	for len(queue) != 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		if c.board[currentPoint.x][currentPoint.y] == 0 {
			continue
		}

		if c.board[currentPoint.x][currentPoint.y] == value {
			if !visited[currentPoint.x][currentPoint.y] {
				visited[currentPoint.x][currentPoint.y] = true
				indexes = append(indexes, currentPoint)
				for _, direction := range directions {
					nextPoint := point{
						currentPoint.x + direction.x,
						currentPoint.y + direction.y,
					}
					queue = append(queue, nextPoint)
				}
			}
		}
	}

	return
}

func (c *Chessboard) remove(minLength int, points []point) bool {
	if len(points) < minLength {
		return false
	}

	for _, p := range points {
		c.board[p.x][p.y] = 0
	}
	return true
}

func (c *Chessboard) fall() {
	matrix := c.board
	row, col := len(matrix), len(matrix[0])
	newMatrix := make([][]int, col)
	for i := range newMatrix {
		newMatrix[i] = make([]int, row)
	}
	for i, row := range matrix {
		for j, val := range row {
			newMatrix[j][len(matrix)-i-1] = val
		}
	}

	for i, row := range newMatrix {
		j := 0
		for ; j < len(row); j++ {
			if row[j] != 0 {
				break
			}
		}
		newMatrix[i] = append(row[j:], make([]int, j)...)
	}

	row, col = len(newMatrix), len(newMatrix[0])
	finalMatrix := make([][]int, col)
	for i := range finalMatrix {
		finalMatrix[i] = make([]int, row)
	}
	for i, row := range newMatrix {
		for j, val := range row {
			finalMatrix[col-j-1][i] = val
		}
	}

	emptyLine := make([]int, len(finalMatrix[0]))
	for i := range emptyLine {
		emptyLine[i] = 0
	}
	c.board = append(finalMatrix[1:], emptyLine)
}

func (c *Chessboard) updateHas() {
	for i := 1; i < len(c.board)-1; i++ {
		for j := 1; j < len(c.board[i])-1; j++ {
			if c.board[i][j] != 0 {
				c.has[point{i, j}] = struct{}{}
			} else {
				delete(c.has, point{i, j})
			}
		}
	}
}

func (c *Chessboard) execute(steps int, minLength int) {
	for i := 0; i < steps; i++ {
		if !c.findAll(minLength) {
			break
		}
		c.fall()
		c.updateHas()
	}
}

func (c *Chessboard) export() [][]int {
	result := make([][]int, len(c.board)-2)
	for i := 1; i < len(c.board)-1; i++ {
		result[i-1] = c.board[i][1 : len(c.board[i])-1]
	}
	return result
}

func NewChessboard(board [][]int) *Chessboard {
	c := Chessboard{
		board: make([][]int, len(board)+2),
		has:   map[point]struct{}{},
	}
	for i := 1; i <= len(board); i++ {
		c.board[i] = make([]int, len(board[i-1])+2)
		copy(c.board[i][1:], board[i-1])
	}

	c.board[0] = make([]int, len(board[0])+2)
	c.board[len(board)+1] = make([]int, len(board[0])+2)
	c.updateHas()

	return &c
}

func solution(n int, k int, input [][]int) [][]int {
	c := NewChessboard(input)
	c.execute(114514, k)
	return c.export()
}
