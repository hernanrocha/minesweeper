package viewmodels

import (
	"time"
)

type CreateGameRequest struct {
	Rows  int `json:"rows"`
	Cols  int `json:"cols"`
	Mines int `json:"mines"`
}

type CreateGameResponse struct {
	ID         string     `json:"id"`
	Rows       int        `json:"rows"`
	Cols       int        `json:"cols"`
	Mines      int        `json:"mines"`
	Board      [][]int    `json:"board"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
}

type RevealRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type FlagRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}
