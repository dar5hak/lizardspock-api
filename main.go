package main

import (
	"fmt"
	"lizardspock/controller"
	_ "lizardspock/docs"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	play := v1.Group("/play")
	{
		play.GET("/1p", controller.Play1p)
		play.GET("/2p", controller.Play2p)
	}

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// @title lizardspock
// @version 1.0
// @license.name Apache 2.0

// @BasePath /api/v1

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := setupRouter()
	err := r.Run(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}
