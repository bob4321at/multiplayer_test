package main

import (
	"main/player"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/AddPlayer", player.AddPlayer)

	r.Run()
}
