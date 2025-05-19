package player

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/textures"
	"main/utils"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerDataForServer struct {
	X  float64
	Y  float64
	ID float64
}

type Player struct {
	Pos            utils.Vec2
	Vel            utils.Vec2
	Speed          float64
	Origonal_Speed float64
	Texture        textures.RenderableTexture
	ID             string
}

func NewPlayer(Pos utils.Vec2, Server_Link string) (player Player) {
	player.Pos = Pos
	player.Speed = 10
	player.Origonal_Speed = player.Speed
	player.Texture = textures.NewTexture("./art/player.png", "")

	json_data, err := json.Marshal(player)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(Server_Link+"/GetIP", "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		panic(err)
	}

	temp_data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(temp_data))

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
}

var Current_Player Player
