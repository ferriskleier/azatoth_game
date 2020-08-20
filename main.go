package main

import (
	"azatoth_game/lib/game"
	"azatoth_game/lib/levels"
	"azatoth_game/lib/player"
)

func main() {

	levels.Intro()
	player.Player()
	game.Game()
}
