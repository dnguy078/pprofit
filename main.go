package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}
