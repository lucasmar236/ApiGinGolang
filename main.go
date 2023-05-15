package main

import (
	"ApiRest/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	services.InitDB()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
