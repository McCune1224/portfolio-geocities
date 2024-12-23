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
	}
}

// Flattens Paddle into format of (x,y,direction)
func (p *Paddle) Flatten() string {
	return fmt.Sprintf("%d,%d,%s", p.coordinate.Y, p.coordinate.X, p.direction.pathing)
}

func (p *Paddle) Move(grid [][]string, nextDirection Direction) {
	p.direction = nextDirection
	if p.direction == DirectionNothing {
		return
	}

	if p.direction == DirectionUp {
		// Check if we're at the top edge or if there's another paddle above
		if p.coordinate.Y <= 0 {
			return
		}
		// Check if there's space to move up (check the space above the paddle's top)
		if grid[p.coordinate.Y-1][p.coordinate.X] == "|" {
			return
		}
		// Move up by updating coordinate
		p.coordinate.Y--
	}

	if p.direction == DirectionDown {
		// Check if we're at the bottom edge or if there's another paddle below
		// Need to check against len-3 because paddle is 3 chars high
		if p.coordinate.Y >= len(grid)-3 {
			return
		}
		// Check if there's space to move down (check the space below the paddle's bottom)
		if grid[p.coordinate.Y+3][p.coordinate.X] == "|" {
			return
		}
		// Move down by updating coordinate
		p.coordinate.Y++
	}
}
