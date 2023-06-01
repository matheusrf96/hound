package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	Port                     = 0
	AbsolutePath             = ""
	RabbitMQConnectionString = ""
)

func Load() {
	gin.SetMode(gin.ReleaseMode)

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil || debug {
		debug = true

		err = godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}

	_, b, _, _ := runtime.Caller(0)
	AbsolutePath = filepath.Join(filepath.Dir(b), "../..")

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 8000
	}

	RabbitMQConnectionString = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	os.Setenv("TZ", "America/Sao_Paulo")
}
