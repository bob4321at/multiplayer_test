package main

import (
	"bufio"
	"fmt"
	"main/player"
	"main/utils"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	player.Current_Player.Update()

	return nil
}

func (g *Game) Draw(s *ebiten.Image) {
	ebitenutil.DebugPrint(s, server_link)

	player.DrawOtherPlayers(s)

	player.Current_Player.Draw(s)
}

func (g *Game) Layout(ow, oh int) (sw, sh int) {
	return 1920, 1080
}

var server_link string

func main() {
	fmt.Print("Enter Server Link: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	server_link = scanner.Text()

	if server_link == "" {
		server_link = "http://localhost:8080"
	}

	player.Current_Player = player.NewPlayer(utils.Vec2{X: 100, Y: 100}, server_link)

	ebiten.SetWindowSize(1920, 1080)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
