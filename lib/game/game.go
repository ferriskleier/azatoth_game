package game

import (
	"azatoth_game/lib/levels"
	"azatoth_game/lib/player"
)

// Game starting
func Game() {

	// block to get player data
	name, race, home, class := player.Strings()
	_ = name
	_ = race
	_ = home
	_ = class

	// first level
	levels.Schnottzslanding(name, class)

	// second level
	levels.SilithusCoast()
}
