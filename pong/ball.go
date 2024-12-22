package ponggame

import "fmt"

type Ball struct {
	coordinate Coordinate
	direction  Direction
}

func (b *Ball) GetShape() [][]string {
	glyph := b.GetSymbol()
	return [][]string{
		{glyph},
	}
}

func (b *Ball) GetCoordinate() Coordinate {
	return b.coordinate
}

func (b *Ball) GetSymbol() string {
	return "*"
}

// Flattens Ball into format of (x,y,direction)
func (b *Ball) Flatten() string {
	return fmt.Sprintf("%d,%d,%s", b.coordinate.X, b.coordinate.Y, b.direction.pathing)
}
