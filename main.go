package main

import (
	"azatoth_game/lib"
	"fmt"
)

func main() {

	lib.Player()

	// block to get player data
	name, race, home, class := lib.Strings()
	fmt.Println(name)
	fmt.Println(race)
	fmt.Println(home)
	fmt.Println(class)
}
