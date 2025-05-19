package main

import (
	"fmt"
	"main/player"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomasen/realip"
)

func main() {
	r := gin.Default()

	r.POST("/GetIP", func(c *gin.Context) {
		ip := realip.RealIP(c.Request)
		fmt.Println(ip)
		c.JSON(http.StatusOK, gin.H{"IP": ip})
	})

	r.POST("/AddPlayer", player.AddPlayer)

	r.Run()
}
