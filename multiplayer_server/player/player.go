package player

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlayerData struct {
	X  float64
	Y  float64
	ID int
}

var players = []PlayerData{}
var total_users = 0

func AddPlayer(c *gin.Context) {
	temp_data_string, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	total_users = total_users + 1
	id := total_users

	var temp_data PlayerData

	if err := json.Unmarshal(temp_data_string, &temp_data); err != nil {
		panic(err)
	}

	temp_data.ID = id

	ReturnIdData := PlayerData{
		0,
		0,
		id,
	}

	c.JSON(http.StatusOK, ReturnIdData)

	fmt.Println(ReturnIdData)
}
