package player

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

type PlayerData struct {
	X  float64
	Y  float64
	ID float64
}

var players = []PlayerData{}

func AddPlayer(c *gin.Context) {
	temp_data_string, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var temp_data PlayerData

	if err := json.Unmarshal(temp_data_string, &temp_data); err != nil {
		panic(err)
	}

	fmt.Println(temp_data)
}
