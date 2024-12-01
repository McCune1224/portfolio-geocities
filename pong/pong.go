package ponggame

import (
	"fmt"
	"strings"
)

const (
	GlobalBoardHeight = 15
	GlobalBoardWidth  = 175
)

// Coordinates are tracked from the top-left most position
type Coordinate struct {
	X int
	Y int
}

type Direction struct {
	pathing string
}

type BoardItem interface {
	GetShape() [][]string
	GetCoordinate() Coordinate
	GetSymbol() string
}

type Ball struct {
	position  Coordinate
	direction Direction
}

func (b *Ball) GetShape() [][]string {
	glyph := b.GetSymbol()
	return [][]string{
		{glyph, glyph},
		{glyph, glyph},
	}
}

func (b *Ball) GetCoordinate() Coordinate {
	return b.GetCoordinate()
}

func (b *Ball) GetSymbol() string {
	return "B"
}

type Paddle struct {
	coordinate Coordinate
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

type BoardState struct {
	Grid [][]string
}

// func (bs *BoardState) Draw(startRow, startCol int, subSection [][]string) error {
func (bs *BoardState) Draw(items ...BoardItem) error {
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

func NewBoard(height int, width int) BoardState {
	grid := [][]string{}
	for range height {
		row := []string{}
		for range width {
			row = append(row, "-")
		}
		grid = append(grid, row)
	}
	board := BoardState{
		Grid: grid,
	}
	playerPaddle := &Paddle{Coordinate{height / 2, width - 15}}
	opponentPaddle := &Paddle{Coordinate{height / 2, 15}}
	// ball := &Ball{position: Coordinate{10, 10}}
	// board.Draw(playerPaddle)
	board.Draw(playerPaddle)
	board.Draw(opponentPaddle)

	return board
}

func NewBoardFromFlattenedBoard(flattenBoard string) BoardState {
	rows := strings.Split(flattenBoard, "\n")
	grid := [][]string{}
	for _, row := range rows {
		grid = append(grid, strings.Split(row, ""))
	}
	return BoardState{
		Grid: grid,
	}
}

func (bs *BoardState) FlattenBoard() string {
	flattenedBoard := ""
	for _, row := range bs.Grid {
		flattenedBoard += strings.Join(row, "") + "\n"
	}
	return flattenedBoard
}
