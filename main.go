package main

import (
	"fmt"
	"hound/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()

	r := gin.New()
	r.Use(gin.Recovery())

	fmt.Println("Hound is running... chasing the rabbit")
	r.Run(fmt.Sprintf(":%d", 8000))
}
