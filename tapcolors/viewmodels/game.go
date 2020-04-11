package viewmodels

import "time"

type CreateGameRequest struct {
}

type CreateGameResponse struct {
	Board     [][]int   `json:"board"`
	Timestamp time.Time `json:"timestamp"`
}

type TapRequest struct {
	Timestamp time.Time `json:"timestamp"`
	Row       int       `json:"row"`
	Col       int       `json:"col"`
	Level     int       `json:"level"`
}

type WebSocketMessage struct {
	Action  string     `json:"action"`
	Payload TapRequest `json:"payload"`
}
