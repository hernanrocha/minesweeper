package models

import (
	"math/rand"
	"time"
)

const (
	CELL_BLANK      = 0
	CELL_BOMB       = -1
	CELL_BLANK_FLAG = -2
	CELL_CLEAR      = -3
	CELL_BOMB_FLAG  = -4
)

type Game struct {
	CreatedAt time.Time
	Rows      int
	Cols      int
	Mines     int
	Board     [][]int
}

var CurrentGame Game

func NewGame() Game {
	game := Game{
		CreatedAt: time.Now(),
		Rows:      5,
		Cols:      5,
		Mines:     3,
	}

	game.Board = make([][]int, game.Rows)

	for i := 0; i < len(game.Board); i++ {
		game.Board[i] = make([]int, game.Cols)
	}

	// Generate random bombs
	game.generateBombs()

	CurrentGame = game

	return game
}

func (g *Game) ToView() Game {
	view := Game{
		CreatedAt: g.CreatedAt,
		Rows:      g.Rows,
		Cols:      g.Cols,
		Mines:     g.Mines,
	}

	view.Board = make([][]int, view.Rows)

	for i := 0; i < len(view.Board); i++ {
		view.Board[i] = make([]int, view.Cols)
	}

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if g.Board[i][j] == CELL_BOMB {
				view.Board[i][j] = CELL_BLANK
			} else if g.Board[i][j] == CELL_BOMB_FLAG {
				view.Board[i][j] = CELL_BLANK_FLAG
			} else {
				view.Board[i][j] = g.Board[i][j]
			}
		}
	}

	return view
}

func (g *Game) RevealCell(row, col int) {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols {
		return
	}

	if g.Board[row][col] == CELL_BLANK {
		count := g.calculateCell(row, col)
		if count == 0 {
			g.Board[row][col] = CELL_CLEAR

			// Reveal neighboor cells
			g.RevealCell(row-1, col-1)
			g.RevealCell(row-1, col)
			g.RevealCell(row-1, col+1)
			g.RevealCell(row, col-1)
			g.RevealCell(row, col+1)
			g.RevealCell(row+1, col-1)
			g.RevealCell(row+1, col)
			g.RevealCell(row+1, col+1)
		} else {
			g.Board[row][col] = count
		}
	}
}

func (g *Game) FlagCell(row, col int) {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols {
		return
	}

	if g.Board[row][col] == CELL_BLANK {
		g.Board[row][col] = CELL_BLANK_FLAG
	} else if g.Board[row][col] == CELL_BLANK_FLAG {
		g.Board[row][col] = CELL_BLANK
	} else if g.Board[row][col] == CELL_BOMB {
		g.Board[row][col] = CELL_BOMB_FLAG
	} else if g.Board[row][col] == CELL_BOMB_FLAG {
		g.Board[row][col] = CELL_BOMB
	}
}

func (g *Game) putBomb(pos int) {
	count := 0
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if count == pos {
				g.Board[i][j] = CELL_BOMB
				return
			}
			if g.Board[i][j] == CELL_BLANK {
				count++
			}
		}
	}
}

func (g *Game) generateBombs() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cells := g.Rows * g.Cols

	for i := 0; i < g.Mines; i++ {
		g.putBomb(r.Intn(cells - i))
	}
}

func (g *Game) isBomb(row, col int) bool {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols {
		return false
	}

	return g.Board[row][col] == CELL_BOMB
}

func (g *Game) calculateCell(row, col int) int {
	bombs := 0

	// Count neighboor bombs
	if g.isBomb(row-1, col-1) {
		bombs++
	}
	if g.isBomb(row-1, col) {
		bombs++
	}
	if g.isBomb(row-1, col+1) {
		bombs++
	}
	if g.isBomb(row, col-1) {
		bombs++
	}
	if g.isBomb(row, col+1) {
		bombs++
	}
	if g.isBomb(row+1, col-1) {
		bombs++
	}
	if g.isBomb(row+1, col) {
		bombs++
	}
	if g.isBomb(row+1, col+1) {
		bombs++
	}

	return bombs
}
