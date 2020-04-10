package models

import (
	"encoding/json"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

const (
	CELL_BLANK         = 0
	CELL_BOMB          = -1
	CELL_BLANK_FLAG    = -2
	CELL_CLEAR         = -3
	CELL_BOMB_FLAG     = -4
	CELL_REVEALED_BOMB = -5
)

type GameStatus string

const (
	STATUS_PLAYING GameStatus = "playing"
	STATUS_LOSE    GameStatus = "lose"
	STATUS_WON     GameStatus = "won"
)

type Game struct {
	ID        string
	CreatedAt time.Time
	Rows      int
	Cols      int
	Mines     int
	Board     [][]int
	Status    GameStatus
}

func NewGame(rows, cols, mines int) *Game {
	game := &Game{
		ID:        generateID(),
		CreatedAt: time.Now(),
		Rows:      rows,
		Cols:      cols,
		Mines:     mines,
		Status:    STATUS_PLAYING,
	}

	// Validate parameters
	if game.Rows < 5 {
		game.Rows = 5
	}
	if game.Rows > 20 {
		game.Rows = 20
	}
	if game.Cols < 5 {
		game.Cols = 5
	}
	if game.Cols > 20 {
		game.Cols = 20
	}
	if game.Mines < 5 {
		game.Mines = 5
	}
	if game.Mines >= game.Rows*game.Cols {
		game.Mines = game.Rows*game.Cols - 1
	}

	// Create Board
	game.Board = make([][]int, game.Rows)

	for i := 0; i < len(game.Board); i++ {
		game.Board[i] = make([]int, game.Cols)
	}

	// Generate random bombs
	game.generateBombs()

	return game
}

func LoadGame(db *redis.Client, id string) (*Game, error) {
	str, err := db.Get(strings.ToUpper(id)).Result()
	if err != nil {
		return nil, err
	}

	var g Game
	if err := json.Unmarshal([]byte(str), &g); err != nil {
		return nil, err
	}
	return &g, nil
}

func (g *Game) ToView() *Game {
	view := &Game{
		ID:        g.ID,
		CreatedAt: g.CreatedAt,
		Rows:      g.Rows,
		Cols:      g.Cols,
		Mines:     g.Mines,
		Status:    g.Status,
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
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols || g.Status != STATUS_PLAYING {
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

		if g.isOver() {
			g.Status = STATUS_WON
		}
	} else if g.Board[row][col] == CELL_BOMB {
		g.Board[row][col] = CELL_REVEALED_BOMB
		g.Status = STATUS_LOSE
	}
}

func (g *Game) FlagCell(row, col int) {
	if row < 0 || row >= g.Rows || col < 0 || col >= g.Cols || g.Status != STATUS_PLAYING {
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

func (g *Game) Save(db *redis.Client) error {
	str, _ := json.Marshal(g)
	return db.Set(strings.ToUpper(g.ID), str, 0).Err()
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

	return g.Board[row][col] == CELL_BOMB || g.Board[row][col] == CELL_BOMB_FLAG
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

func (g *Game) isOver() bool {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			if g.Board[i][j] == CELL_BLANK || g.Board[i][j] == CELL_BLANK_FLAG {
				return false
			}
		}
	}
	return true
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateID() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
