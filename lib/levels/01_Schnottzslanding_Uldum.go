package levels

import (
	"azatoth_game/lib/player"
	"fmt"
	"time"
)

// Schnottzslanding level
func Schnottzslanding() {

	// block to get player data
	name, race, home, class := player.Strings()
	_ = race
	_ = home

	// Intro
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
	fmt.Printf("=================================\n" +
		"    Chapter 1: The Shattering\n" +
		"=================================\n\n" +
		"Welcome " + name + "\n" +
		"You're on the 'Dry Seaheart', a ship full of members of\n" +
		"the league of explorers. They measured activity in front\n" +
		"of Ahn'Qiraj, the temple and prison of C'thun. C'thun is an\n" +
		"old god who was defeated years ago, but never died. To\n" +
		"this day he tries to break his chains and both the league\n" +
		"of explorers and guards of Silithus want to know what\n" +
		"happens inside of Ahn'Qiraj.\n" +
		"You help this team of explorers to get a safe access into\n" +
		"Ahn'Qiraj\n\n" +
		"=======================\n" +
		"Press enter to continue\n" +
		"=======================\n\n\n")
	fmt.Scanln()

	// Part 1
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
	fmt.Printf("============================================\n" +
		"Location: 'Dry Seaheart', Schnottz's Landing\n" +
		"============================================\n\n\n" +
		"Harrison Jones:\n" +
		"===============\n" +
		"Welcome, " + name + "!\n" +
		"I am Harrison Jones, archeological leader of the league\n" +
		"of explorers. I hope you got everything with you, you need.\n" +
		"We will start in a few minutes, leaving Schnottz's landing\n" +
		"and heading to the hilly, southern coast of Silithus.\n" +
		"There we will moor the ship and preapare in the bay we\n" +
		"located earlier. You can rest until then, it takes us a\n" +
		"while to arrive at the location. I count on you, " + class + ".\n\n" +
		"(press enter to sleep)\n")
	fmt.Scanln()

	// Part 2
	for i := 0; i < 5; i++ {
		time.Sleep(1100 * time.Millisecond)
		fmt.Println("  ...")
	}
	fmt.Printf("\n  , ., , ., . ,. ,. ,. , ., , .\n" +
		"~x~X~X`X~x~x`x~X~XX~X~X~x`X~X~X~X\n" +
		"  ' `' `'`''` '` '`' `'` '`' '`'`\n\n")
	time.Sleep(900 * time.Millisecond)
	fmt.Scanln()
	fmt.Printf("" +
		"Harrison Jones:\n" +
		"===============\n" +
		"What happened?\n" +
		"It seems we crashed. But why? The coast should be one\n" +
		"hour away.\n")
	fmt.Scanln()
	fmt.Printf("*Harrison and you run up to the deck, to see what happened*\n")
	fmt.Scanln()
	fmt.Printf("*A heavy storm slams on the ship and waves crash on it's\n" +
		"sides. You see the upper deck full of people and tauren\n" +
		"shamans trying to keep the water from flooding the ship*\n")
	fmt.Scanln()
	fmt.Printf("Harrison Jones:\n" +
		"===============\n" +
		"Well, we obviously crashed into the cliffs of Silithus, but\n" +
		"as I already said we should be far away from the coast.\n")
	fmt.Scanln()
	fmt.Printf("*The captain of the ship, Thasarah Vineglade, walks toward\n" +
		"you and Harrison*\n")
	fmt.Scanln()
}
