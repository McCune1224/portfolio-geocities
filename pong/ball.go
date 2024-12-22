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

// Update the ball's position based on its direction
func (b *Ball) Move(grid [][]string, playerPaddle Paddle, opponentPaddle Paddle) {
	if b.direction == DirectionUpRight {
		if b.coordinate.X == 0 || b.coordinate.Y == len(grid[0])-1 || grid[b.coordinate.X-1][b.coordinate.Y+1] == "|" {
			b.direction = DirectionDownRight
			b.coordinate.X++
			b.coordinate.Y++
		} else if b.coordinate.Y == len(grid[0])-1 {
			b.direction = DirectionUpLeft
			b.coordinate.Y--
		} else {
			b.coordinate.X--
			b.coordinate.Y++
		}
	}

	if b.direction == DirectionUpLeft {
		if b.coordinate.X == 0 || b.coordinate.Y == 0 || grid[b.coordinate.X-1][b.coordinate.Y-1] == "|" {
			b.direction = DirectionDownLeft
			b.coordinate.X++
			b.coordinate.Y--
		} else if b.coordinate.Y == 0 {
			b.direction = DirectionUpRight
			b.coordinate.Y++
		} else {
			b.coordinate.X--
			b.coordinate.Y--
		}
	}

	if b.direction == DirectionDownRight {
		if b.coordinate.X == len(grid)-1 || b.coordinate.Y == len(grid[0])-1 || grid[b.coordinate.X+1][b.coordinate.Y+1] == "|" {
			b.direction = DirectionUpRight
			b.coordinate.X--
			b.coordinate.Y++
		} else if b.coordinate.Y == len(grid[0])-1 {
			b.direction = DirectionDownLeft
			b.coordinate.Y--
		} else {
			b.coordinate.X++
			b.coordinate.Y++
		}
	}

	if b.direction == DirectionDownLeft {
		if b.coordinate.X == len(grid)-1 || b.coordinate.Y == 0 || grid[b.coordinate.X+1][b.coordinate.Y-1] == "|" {
			b.direction = DirectionUpLeft
			b.coordinate.X--
			b.coordinate.Y--
		} else if b.coordinate.Y == 0 {
			b.direction = DirectionDownRight
			b.coordinate.Y++
		} else {
			b.coordinate.X++
			b.coordinate.Y--
		}
	}
}
