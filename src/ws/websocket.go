package ws

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func WsEndpoint(c *gin.Context) {
	ws, err := websocket.Accept(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}
	defer ws.Close(websocket.StatusInternalError, "the sky is falling")

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	messageType, message, err := ws.Read(ctx)
	if err != nil {
		log.Println(err)
	}

	log.Printf("received: %s: %s", messageType.String(), string(message))

	ws.Close(websocket.StatusNormalClosure, "")
}
