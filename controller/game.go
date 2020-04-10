package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/hernanrocha/minesweeper/models"
	"github.com/hernanrocha/minesweeper/viewmodels"
)

// GameController ...
type GameController struct {
	db *gorm.DB
}

// NewGameController ...
func NewGameController() *GameController {
	return &GameController{
		// db: models.GetDB(),
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

	game := models.NewGame()
	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		Rows:      game.Rows,
		Cols:      game.Cols,
		Mines:     game.Mines,
		Board:     game.Board,
		CreatedAt: game.CreatedAt,
		Status:    string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}

// RevealCell godoc
// @Summary Reveal Cell
// @Description Reveal cell in board
// @Tags Game
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

	game := models.CurrentGame
	game.RevealCell(json.Row, json.Col)
	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		Rows:      game.Rows,
		Cols:      game.Cols,
		Mines:     game.Mines,
		Board:     game.Board,
		CreatedAt: game.CreatedAt,
		Status:    string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}

// FlagCell godoc
// @Summary Flag Cell
// @Description Flag cell in board
// @Tags Game
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

	game := models.CurrentGame
	game.FlagCell(json.Row, json.Col)
	game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		Rows:      game.Rows,
		Cols:      game.Cols,
		Mines:     game.Mines,
		Board:     game.Board,
		CreatedAt: game.CreatedAt,
		Status:    string(game.Status),
	}

	ctx.JSON(http.StatusOK, response)
}
