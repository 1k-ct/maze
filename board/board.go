package board

// Size  盤の大きさ
type Size struct {
	Width  int
	Height int
}

type Map struct {
	Maze [][]int
}

// New some
func New(size, wall, aisle int) *Size {
	s := &Size{Width: size, Height: size}
	return s
}
func (s Size) MakeField() *Map {
	m := make([][]int, s.Height)
	for i := 0; i < s.Height; i++ {
		m[i] = make([]int, s.Width)
	}
	ma := &Map{Maze: m}
	s.createWall(ma)
	return ma
}

func (s Size) createWall(m *Map) {
	wall := 1
	for i := 0; i < s.Height; i++ {
		for i := 0; i < s.Width; i++ {
			m.Maze[0][i] = wall
			m.Maze[s.Height-1][i] = wall
		}
		m.Maze[i][0] = wall
		m.Maze[i][s.Width-1] = wall
	}
}
