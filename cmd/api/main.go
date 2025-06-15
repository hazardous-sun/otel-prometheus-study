package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
			"status":  http.StatusOK,
		})
	})

	err := server.Run(":8000")

	if err != nil {
		log.Fatal(err)
		return
	}
}
