package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hernanrocha/minesweeper/tapcolors/models"
	"github.com/hernanrocha/minesweeper/tapcolors/viewmodels"
)

// GameController ...
type GameController struct {
}

// NewGameController ...
func NewGameController() *GameController {
	return &GameController{}
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

	game := models.NewRandom(4)

	// log.Println(game.FindSolution())
	// game.Save(c.db)
	// game = game.ToView()

	response := &viewmodels.CreateGameResponse{
		// ID:         game.ID,
		Board: game.Board,
	}

	ctx.JSON(http.StatusOK, response)
}
