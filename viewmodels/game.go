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
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type RevealRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type FlagRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
