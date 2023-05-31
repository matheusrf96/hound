package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	fmt.Println("Hound is running... behing a rabbit")
	r.Run(fmt.Sprintf(":%d", 8000))
}
