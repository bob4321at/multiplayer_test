package player

import (
	"bytes"
	"encoding/json"
	"io"
	"main/utils"
	"net/http"

	"github.com/bob4321at/textures"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerDataForServer struct {
	X  float64
	Y  float64
	ID int
}

type Player struct {
	Pos            utils.Vec2
	Vel            utils.Vec2
	Speed          float64
	Origonal_Speed float64
	Texture        textures.RenderableTexture
	ID             int
	ServerLink     string
}

func NewPlayer(Pos utils.Vec2, Server_Link string) (player Player) {
	player.Pos = Pos
	player.Speed = 10
	player.Origonal_Speed = player.Speed
	player.Texture = textures.NewTexture("./art/player.png", "")

	json_data, err := json.Marshal(player.GetPlayerServerData())
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(Server_Link+"/AddPlayer", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		panic(err)
	}

	temp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var get_id_data PlayerDataForServer

	if err := json.Unmarshal(temp_data, &get_id_data); err != nil {
		panic(err)
	}

	player.ID = get_id_data.ID

	player.ServerLink = Server_Link

	return player
}

func (player *Player) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(player.Pos.X, player.Pos.Y)

	player.Texture.Draw(screen, &op)
}

func (player *Player) Update() {
	player.Speed = player.Origonal_Speed

	player.Vel = utils.Vec2{X: 0, Y: 0}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.Vel.Y = -1
		player.Speed /= 1.5
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		player.Vel.Y = 1
		player.Speed /= 1.5
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.Vel.X = 1
		player.Speed /= 1.5
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.Vel.X = -1
		player.Speed /= 1.5
	}

	player.Pos.X += player.Vel.X * player.Speed
	player.Pos.Y += player.Vel.Y * player.Speed

	json_data, err := json.Marshal(player.GetPlayerServerData())
	if err != nil {
		panic(err)
	}

	http.Post(player.ServerLink+"/MovePlayer", "application/json", bytes.NewBuffer(json_data))
}

func (player *Player) GetPlayerServerData() (playerData PlayerDataForServer) {
	playerData.X = player.Pos.X
	playerData.Y = player.Pos.Y
	playerData.ID = player.ID

	return playerData
}

var other_player_texture = textures.NewTexture("./art/player.png", "")

func DrawOtherPlayers(screen *ebiten.Image) {
	json_data, err := json.Marshal(Current_Player.GetPlayerServerData())
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(Current_Player.ServerLink+"/GetOtherPlayers", "application/json", bytes.NewBuffer(json_data))

	var temp_data []PlayerDataForServer

	data_to_read, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data_to_read, &temp_data); err != nil {
		panic(err)
	}

	op := ebiten.DrawImageOptions{}
	for _, user := range temp_data {
		op.GeoM.Reset()
		op.GeoM.Translate(user.X, user.Y)
		other_player_texture.Draw(screen, &op)
	}
}

var Current_Player Player
