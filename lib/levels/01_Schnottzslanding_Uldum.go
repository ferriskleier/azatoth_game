package levels

import (
	"azatoth_game/lib/player"
	"fmt"
)

// Schnottzslanding level
func Schnottzslanding() {

	// block to get player data
	name, race, home, class := player.Strings()
	_ = name
	_ = race
	_ = home
	_ = class

	fmt.Println("Test")
}
