package ponggame

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

const (
	GlobalBoardHeight = 10
	GlobalBoardWidth  = 110
)

var (
	DirectionUp        = Direction{"U"}
	DirectionDown      = Direction{"D"}
	DirectionLeft      = Direction{"L"}
	DirectionRight     = Direction{"R"}
	DirectionUpLeft    = Direction{"UL"}
	DirectionUpRight   = Direction{"UR"}
	DirectionDownLeft  = Direction{"DL"}
	DirectionDownRight = Direction{"DR"}
	DirectionNothing   = Direction{"N"}
)

// Coordinates are tracked from the top-left most position
type Coordinate struct {
	X int
	Y int
}

type Direction struct {
	// Pathing is a string that represents the direction of the ball. UR = Up Right, DL = Down Left, ...
	pathing string
}

type BoardItem interface {
	GetShape() [][]string
	GetCoordinate() Coordinate
	GetSymbol() string
}

type BoardState struct {
	Grid            [][]string
	BackgroundGlyph string
	ball            *Ball
	playerPaddle    *Paddle
	opponentPaddle  *Paddle
}

// Draws the Paddles and Ball onto the board
func (bs *BoardState) DrawItems() error {
	items := []BoardItem{bs.ball, bs.playerPaddle, bs.opponentPaddle}
	//Reset the board
	for i := range bs.Grid {
		for j := range bs.Grid[i] {
			bs.Grid[i][j] = bs.BackgroundGlyph
		}
	}
	for _, item := range items {
		startRow := item.GetCoordinate().X
		startCol := item.GetCoordinate().Y
		subSection := item.GetShape()
		// Check if subsection fits within grid bounds
		if startRow < 0 || startCol < 0 {
			return fmt.Errorf("starting position (%d,%d) out of bounds", startRow, startCol)
		}

		if len(subSection) == 0 || len(subSection[0]) == 0 {
			return fmt.Errorf("subsection cannot be empty")
		}

		endRow := startRow + len(subSection)
		endCol := startCol + len(subSection[0])

		if endRow > len(bs.Grid) || endCol > len(bs.Grid[0]) {
			return fmt.Errorf("subsection exceeds grid bounds")
		}

		// Update the subsection
		for i := range subSection {
			for j := range subSection[i] {
				bs.Grid[startRow+i][startCol+j] = subSection[i][j]
			}
		}
	}
	return nil
}

// Move all board items in the direction they are facing (redirecting if they hit a wall) and redraw the board
func (bs *BoardState) UpdateBoard(newPlayerDirection Direction) error {
	// Ball Logic
	bs.ball.Move(bs.Grid, *bs.playerPaddle, *bs.opponentPaddle)
	bs.playerPaddle.Move(bs.Grid, newPlayerDirection)
	bs.playerPaddle.Move(bs.Grid, newPlayerDirection)

	// Paddle Logic

	return bs.DrawItems()
}

func NewBoard(height int, width int) BoardState {
	backgroundGlyph := " - "
	grid := [][]string{}
	for range height {
		row := []string{}
		for range width {
			// FIXME: Need to add a whitespace because the ball glyph makes the board slanted on the right side of the ball
			row = append(row, backgroundGlyph)
		}
		grid = append(grid, row)
	}

	startingBallDirections := []Direction{DirectionUpRight, DirectionDownRight, DirectionUpLeft, DirectionDownLeft}

	board := BoardState{
		Grid:            grid,
		playerPaddle:    &Paddle{coordinate: Coordinate{(height / 2) - 3, width - 15}, direction: DirectionNothing},
		opponentPaddle:  &Paddle{coordinate: Coordinate{(height / 2) - 3, 15}, direction: DirectionNothing},
		ball:            &Ball{coordinate: Coordinate{height / 2, width / 2}, direction: startingBallDirections[rand.IntN(len(startingBallDirections))]},
		BackgroundGlyph: backgroundGlyph,
	}
	board.DrawItems()

	return board
}

func NewBoardFromFlattenedBoard(flattenBoard string) BoardState {
	split := strings.Split(flattenBoard, "_")
	boardStr := split[0]
	ballStr := split[1]
	playerStr := split[2]
	opponentStr := split[3]

	boardSplit := strings.Split(boardStr, ",")
	bX, _ := strconv.Atoi(boardSplit[0])
	bY, _ := strconv.Atoi(boardSplit[1])

	ballSplit := strings.Split(ballStr, ",")
	ballX, _ := strconv.Atoi(ballSplit[0])
	ballY, _ := strconv.Atoi(ballSplit[1])
	ballDirection := Direction{ballSplit[2]}
	ball := Ball{
		coordinate: Coordinate{ballX, ballY},
		direction:  ballDirection,
	}

	playerSplit := strings.Split(playerStr, ",")
	playerX, _ := strconv.Atoi(playerSplit[0])
	playerY, _ := strconv.Atoi(playerSplit[1])
	playerDirection := Direction{playerSplit[2]}
	player := Paddle{coordinate: Coordinate{playerX, playerY}, direction: playerDirection}

	opponentSplit := strings.Split(opponentStr, ",")
	opponentX, _ := strconv.Atoi(opponentSplit[0])
	opponentY, _ := strconv.Atoi(opponentSplit[1])
	opponentDirection := Direction{opponentSplit[2]}
	opponent := Paddle{coordinate: Coordinate{opponentX, opponentY}, direction: opponentDirection}

	grid := [][]string{}

	for range bX {
		row := []string{}
		for range bY {
			row = append(row, "- ")
		}
		grid = append(grid, row)
	}

	board := BoardState{
		BackgroundGlyph: "- ",
		Grid:            grid,
		playerPaddle:    &player,
		opponentPaddle:  &opponent,
		ball:            &ball,
	}
	board.DrawItems()
	return board
}

// Flatten puts the board into a string format of (grid, ball, player, opponent)
func (bs *BoardState) FlattenBoard() string {
	templatestr := "%s_%s_%s_%s"
	gridStr := fmt.Sprintf("%d,%d", len(bs.Grid), len(bs.Grid[0]))
	ballStr := bs.ball.Flatten()
	opponentStr := bs.opponentPaddle.Flatten()
	playerStr := bs.playerPaddle.Flatten()

	return fmt.Sprintf(templatestr, gridStr, ballStr, playerStr, opponentStr)
}

func (bs *BoardState) GetBall() Ball {
	var ball Ball
	for x, row := range bs.Grid {
		for y, cell := range row {
			if cell == ball.GetSymbol() {
				return Ball{
					coordinate: Coordinate{x, y},
				}
			}
		}
	}
	return ball
}

func (bs *BoardState) GetPlayerPaddle() Paddle {
	var paddle Paddle
	for x, row := range bs.Grid {
		for y, cell := range row {
			if cell == "|" {
				return Paddle{
					coordinate: Coordinate{x, y},
				}
			}
		}
	}
	return paddle
}

func (bs *BoardState) GetOpponentPaddle() Paddle {
	var paddle Paddle
	for x, row := range bs.Grid {
		for y, cell := range row {
			if cell == "|" {
				return Paddle{
					coordinate: Coordinate{x, y},
				}
			}
		}
	}
	return paddle
}
