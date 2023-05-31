package main

import (
	"fmt"
	"hound/src/config"
	"hound/src/ws"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Ok"}) }) // Readiness Probe
	r.GET("/ws", ws.WsEndpoint)
}

func main() {
	config.Load()

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	setupRoutes(r)

	fmt.Println("Hound is running... chasing the rabbit")
	r.Run(fmt.Sprintf(":%d", config.Port))
}
