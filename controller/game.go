package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-redis/redis/v7"
	"github.com/hernanrocha/minesweeper/models"
	"github.com/hernanrocha/minesweeper/viewmodels"
)

// GameController ...
type GameController struct {
	db *redis.Client
}

// NewGameController ...
func NewGameController(db *redis.Client) *GameController {
	return &GameController{
		db: db,
	}
}

// CreateGame godoc
// @Summary Create Game
// @Description Create Game in database
// @Tags Game
// @Param user body viewmodels.CreateGameRequest true "Game Data"
// @Produce  json
// @Success 200 {object} viewmodels.CreateGameResponse
// @Router /api/v1/game [post]
func (c *GameController) CreateGame(ctx *gin.Context) {
	var json viewmodels.CreateGameRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := models.NewGame(json.Rows, json.Cols, json.Mines)
	game.Save(c.db)
	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		ID:         game.ID,
		Rows:       game.Rows,
		Cols:       game.Cols,
		Mines:      game.Mines,
		Board:      game.Board,
		CreatedAt:  game.CreatedAt,
		FinishedAt: game.FinishedAt,
		Status:     string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}

// GetGame godoc
// @Summary Get Game
// @Description Get Game from database
// @Tags Game
// @Param id path int true "Game ID"
// @Produce  json
// @Success 200 {object} viewmodels.CreateGameResponse
// @Router /api/v1/game/{id} [get]
func (c *GameController) GetGame(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	game, err := models.LoadGame(c.db, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		ID:         game.ID,
		Rows:       game.Rows,
		Cols:       game.Cols,
		Mines:      game.Mines,
		Board:      game.Board,
		CreatedAt:  game.CreatedAt,
		FinishedAt: game.FinishedAt,
		Status:     string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}

// RevealCell godoc
// @Summary Reveal Cell
// @Description Reveal cell in board
// @Tags Game
// @Param id path int true "Game ID"
// @Param user body viewmodels.RevealRequest true "Cell Data"
// @Produce  json
// @Success 200 {object} viewmodels.CreateGameResponse
// @Router /api/v1/game/{id}/reveal [post]
func (c *GameController) RevealCell(ctx *gin.Context) {
	var json viewmodels.RevealRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Params.ByName("id")
	game, err := models.LoadGame(c.db, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	game.RevealCell(json.Row, json.Col)
	game.Save(c.db)
	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		ID:         game.ID,
		Rows:       game.Rows,
		Cols:       game.Cols,
		Mines:      game.Mines,
		Board:      game.Board,
		CreatedAt:  game.CreatedAt,
		FinishedAt: game.FinishedAt,
		Status:     string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}

// FlagCell godoc
// @Summary Flag Cell
// @Description Flag cell in board
// @Tags Game
// @Param id path int true "Game ID"
// @Param user body viewmodels.FlagRequest true "Cell Data"
// @Produce  json
// @Success 200 {object} viewmodels.CreateGameResponse
// @Router /api/v1/game/{id}/flag [post]
func (c *GameController) FlagCell(ctx *gin.Context) {
	var json viewmodels.RevealRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Params.ByName("id")
	game, err := models.LoadGame(c.db, id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	game.FlagCell(json.Row, json.Col)
	game.Save(c.db)
	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		ID:         game.ID,
		Rows:       game.Rows,
		Cols:       game.Cols,
		Mines:      game.Mines,
		Board:      game.Board,
		CreatedAt:  game.CreatedAt,
		FinishedAt: game.FinishedAt,
		Status:     string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}
