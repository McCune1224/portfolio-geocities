package ponggame

import "fmt"

type Paddle struct {
	coordinate Coordinate
	direction  Direction
}

// GetCoordinate implements BoardItem.
func (p *Paddle) GetCoordinate() Coordinate {
	return p.coordinate
}

func (p *Paddle) GetSymbol() string {
	return "|"
}

func (p *Paddle) GetShape() [][]string {
	glyph := p.GetSymbol()
	return [][]string{
		{glyph},
		{glyph},
		{glyph},
		{glyph},
		{glyph},
	}
}

// Flattens Paddle into format of (x,y,direction)
func (p *Paddle) Flatten() string {
	return fmt.Sprintf("%d,%d,%s", p.coordinate.X, p.coordinate.Y, p.direction.pathing)
}
