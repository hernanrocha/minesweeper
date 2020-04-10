package models

import "time"

const (
	CELL_BLANK = 0
	CELL_BOMB  = -1
	CELL_FLAG  = -2
)

type Game struct {
	CreatedAt time.Time
	Rows      int
	Cols      int
	Mines     int
	Board     [][]int
}

var currentGame Game

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

	// TODO: Generate random
	game.Board[0][0] = CELL_BOMB
	game.Board[2][1] = CELL_BOMB
	game.Board[4][4] = CELL_BOMB

	currentGame = game

	return game
}
