package main

import (
	"fmt"
	"net/http"
	"sub/src/config"
	"sub/src/workers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "Ok"}) }) // Readiness Probe
}

func main() {
	config.Load()

	r := gin.New()
	r.Use(gin.Recovery())

	setupRoutes(r)

	workers.New(workers.Counter, 3)
	workers.New(workers.GetRabbitMQData, 5)

	fmt.Println("Hound is running... chasing the rabbit")
	fmt.Printf(":%d\n", config.Port)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
