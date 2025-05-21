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
		temp_data.X,
		temp_data.Y,
		id,
	}

	c.JSON(http.StatusOK, ReturnIdData)

	fmt.Println(ReturnIdData)

	players = append(players, ReturnIdData)
}

func MovePlayer(c *gin.Context) {
	temp_data_string, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var temp_data PlayerData

	if err := json.Unmarshal(temp_data_string, &temp_data); err != nil {
		panic(err)
	}

	for i, user := range players {
		if user.ID == temp_data.ID {
			players[i] = temp_data
		}
	}
}

func GetOtherPlayers(c *gin.Context) {
	temp_data_string, err := io.ReadAll(c.Request.Body)
	if err != nil {
		panic(err)
	}

	var temp_data PlayerData

	if err := json.Unmarshal(temp_data_string, &temp_data); err != nil {
		panic(err)
	}

	players_to_send := []PlayerData{}

	for _, user := range players {
		if temp_data.ID != user.ID {
			// temp_user := PlayerData{user.X, user.Y, -1}
			players_to_send = append(players_to_send, user)
		}
	}

	c.JSON(http.StatusOK, players_to_send)
}
