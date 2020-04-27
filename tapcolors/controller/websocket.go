package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/hernanrocha/minesweeper/tapcolors/models"
	"github.com/hernanrocha/minesweeper/tapcolors/viewmodels"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WebSocketController ...
type WebSocketController struct {
	// game    *models.Game
	// clients map[string]*websocket.Conn
}

// NewWebSocketController ...
func NewWebSocketController() *WebSocketController {
	return &WebSocketController{
		// clients: make(map[string]*websocket.Conn),
	}
}

func (c *WebSocketController) WebSocket(ctx *gin.Context) {
	log.Println("Creating new WebSocket")
	newconn, err := wsupgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %s\n", err)
		return
	}

	// handler := handler.NewWebSocketMessageHandler(conn)
	// c.hub.AddClient(handler)
	// log.Println("Adding client...")
	// c.clients[newconn.RemoteAddr().String()] = newconn

	go func(conn *websocket.Conn) {
		game := models.NewRandom(4)
		response := &viewmodels.CreateGameResponse{
			Timestamp: time.Now(),
			Board:     game.Board,
		}

		conn.WriteJSON(response)
		for {
			// Read Message
			_, obj, err := conn.ReadMessage()

			// Handle Message
			var msg viewmodels.WebSocketMessage
			json.Unmarshal(obj, &msg)
			game = Handle(&msg, game)

			// Send Response
			response := &viewmodels.CreateGameResponse{
				Timestamp: msg.Payload.Timestamp,
				Board:     game.Board,
			}
			conn.WriteJSON(response)

			if err != nil {
				log.Println("ERROR ON WEBSOCKET: CLOSING...")
				log.Println(err)
				// c.hub.RemoveClient(handler)
				// delete(c.clients, conn.RemoteAddr().String())
				return
			}
		}
	}(newconn)
}

func Handle(msg *viewmodels.WebSocketMessage, game *models.Game) *models.Game {
	// log.Println("Received: ", msg.Action)
	switch msg.Action {
	case "newgame":
		game = models.NewRandom(msg.Payload.Level)
		log.Println("New Game Created", msg.Payload.Level)
	case "tap":
		game.Tap(msg.Payload.Row, msg.Payload.Col)
	case "solve":
		if game.Level <= 5 {
			log.Println(game.FindSolution())
		}
	default:
		log.Println("Unknown Command")
	}
	return game
}
