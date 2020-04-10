package viewmodels

import (
	"time"
)

type CreateGameRequest struct{}

type CreateGameResponse struct {
	Rows      int       `json:"rows"`
	Cols      int       `json:"cols"`
	Mines     int       `json:"mines"`
	Board     [][]int   `json:"board"`
	CreatedAt time.Time `json:"created_at"`
}
