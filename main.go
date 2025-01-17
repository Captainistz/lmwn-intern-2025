package main

import (
	"log"
	"os"

	"github.com/Captainistz/lmwn-intern-2025/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("[ENV] Unable to load .env")
	}
}

func main() {
	r := gin.Default()
	routers.SetupRoutes(r)

	PORT, ok := os.LookupEnv("PORT")
	if !ok {
		PORT = "8080"
	}

	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatalln("Failed to run server: ", err.Error())
	}
}
