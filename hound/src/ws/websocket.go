package ws

import (
	"context"
	"hound/src/controllers"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func WsEndpoint(c *gin.Context) {
	ws, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		OriginPatterns: []string{"*"},
	})
	if err != nil {
		log.Println(err)
	}
	defer ws.Close(websocket.StatusInternalError, "the sky is falling")

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*10)
	defer cancel()

	_, message, err := ws.Read(ctx)
	if err != nil {
		log.Println(err)
	}

	originIp := c.Request.Header.Get("X-Forwarded-For")
	if originIp == "" {
		originIp = strings.Split(c.Request.RemoteAddr, ":")[0]
	}

	ch := controllers.HandleData(message, originIp)
	cData := <-ch

	if cData != "" {
		log.Println(cData)
	}

	ws.Close(websocket.StatusNormalClosure, "WebSocket executed succesfully")
}
