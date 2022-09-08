package lib

import (
	"errors"
	"fmt"
)

type GameBoard struct {
	Grid Grid
}

type Grid [GRID_HEIGHT][GRID_WIDTH]GridCell

type GridCell uint8

const (
	GRID_WIDTH  = 7
	GRID_HEIGHT = 6
)

const (
	BLANK GridCell = iota
	YELLOW
	RED
)

type Player uint8

const (
	PLAYER_ONE Player = 1
	PLAYER_TWO Player = 2
)

func NewGameBoard() GameBoard {
	var grid Grid

	for i := 0; i < GRID_HEIGHT; i++ {
		for j := 0; j < GRID_WIDTH; j++ {
			grid[i][j] = BLANK
		}
	}

	return GameBoard{
		grid,
	}
}

var ErrColOutOfRange = errors.New("column value out of range")
var ErrFilledBoard = errors.New("column already filled to the top")

func (g *GameBoard) DropDisc(column uint8, player Player) error {
	if column >= GRID_WIDTH {
		return ErrColOutOfRange
	}

	for i := 0; i < GRID_HEIGHT; i++ {
		is_bottom_cell := i == 5
		is_cell_filled := g.Grid[i][column] != BLANK

		if is_cell_filled {
			i -= 1
		}

		if is_cell_filled || is_bottom_cell {
			if i == -1 {
				return ErrColOutOfRange
			}

			switch player {
			case PLAYER_ONE:
				g.Grid[i][column] = YELLOW
			case PLAYER_TWO:
				g.Grid[i][column] = RED
			}

			break
		}
	}

	return nil
}

func (g *GameBoard) Display() {
	for i := 0; i < GRID_HEIGHT; i++ {
		for j := 0; j < GRID_WIDTH; j++ {
			var symbol rune

			switch g.Grid[i][j] {
			case BLANK:
				symbol = '⬜'
			case YELLOW:
				symbol = '🟡'
			case RED:
				symbol = '🔴'
			}

			fmt.Printf("%v", string(symbol))
		}
		fmt.Printf("\n")
	}
}
