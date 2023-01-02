package canvas

type Position struct {
	x1, y1 int
	x2, y2 int
}

func NewPosition(x1, y1, x2, y2 int) *Position {
	return &Position{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
	}
}

func (pos *Position) Set(x1, y1, x2, y2 int) {
	pos.x1 = x1
	pos.x2 = x2
	pos.y1 = y1
	pos.y2 = y2
}
