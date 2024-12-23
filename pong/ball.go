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
	return fmt.Sprintf("%d,%d,%s", b.coordinate.Y, b.coordinate.X, b.direction.pathing)
}

// Update the ball's position based on its direction, will return a string of "L" or "R" if the ball has scored
func (b *Ball) Move(grid [][]string, playerPaddle Paddle, opponentPaddle Paddle) string {

	// fmt.Printf("X: %d Y: %d\n", b.coordinate.Y, b.coordinate.X)
	if b.direction == DirectionUpRight {
		if b.coordinate.Y == 0 || b.coordinate.X == len(grid[0])-1 || grid[b.coordinate.Y-1][b.coordinate.X+1] == playerPaddle.GetSymbol() {
			b.direction = DirectionDownRight
			b.coordinate.Y++
			b.coordinate.X++
		} else if b.coordinate.X == len(grid[0])-1 {
			b.direction = DirectionUpLeft
			b.coordinate.X--
		} else {
			b.coordinate.Y--
			b.coordinate.X++
		}
		return ""
	}

	if b.direction == DirectionUpLeft {
		if b.coordinate.Y == 0 || b.coordinate.X == 0 || grid[b.coordinate.Y-1][b.coordinate.X-1] == playerPaddle.GetSymbol() {
			b.direction = DirectionDownLeft
			b.coordinate.Y++
			b.coordinate.X--
		} else if b.coordinate.X == 0 {
			b.direction = DirectionUpRight
			b.coordinate.X++
		} else {
			b.coordinate.Y--
			b.coordinate.X--
			return ""
		}
	}

	if b.direction == DirectionDownRight {
		if b.coordinate.Y == len(grid)-1 || b.coordinate.X == len(grid[0])-1 || grid[b.coordinate.Y+1][b.coordinate.X+1] == playerPaddle.GetSymbol() {
			b.direction = DirectionUpRight
			b.coordinate.Y--
			b.coordinate.X++
		} else if b.coordinate.X == len(grid[0])-1 {
			b.direction = DirectionDownLeft
			b.coordinate.X--
		} else {
			b.coordinate.Y++
			b.coordinate.X++
			return ""
		}
	}

	if b.direction == DirectionDownLeft {
		if b.coordinate.Y == len(grid)-1 || b.coordinate.X == 0 || grid[b.coordinate.Y+1][b.coordinate.X-1] == playerPaddle.GetSymbol() {
			b.direction = DirectionUpLeft
			b.coordinate.Y--
			b.coordinate.X--
		} else if b.coordinate.X == 0 {
			b.direction = DirectionDownRight
			b.coordinate.X++
		} else {
			b.coordinate.Y++
			b.coordinate.X--
			return ""
		}
	}
	return ""
}
