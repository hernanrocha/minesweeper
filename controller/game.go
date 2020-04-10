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

	response := &viewmodels.CreateGameResponse{
		Rows:      game.Rows,
		Cols:      game.Cols,
		Mines:     game.Mines,
		Board:     game.Board,
		CreatedAt: game.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}
