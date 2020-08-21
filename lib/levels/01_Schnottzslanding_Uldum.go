package levels

import (
	"fmt"
	"time"
)

// Schnottzslanding level 1
func Schnottzslanding(name string, class string) {

	// Intro
	SchnottzslandingIntro(name)

	// Part 1
	Schnottzslanding01(name, class)

	// Part 2
	Schnottzslanding02(name)
}

// SchnottzslandingIntro do
func SchnottzslandingIntro(name string) {

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
}

// Schnottzslanding01 do
func Schnottzslanding01(name string, class string) {
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
}

// Schnottzslanding02 do
func Schnottzslanding02(name string) {
	for i := 0; i < 5; i++ {
		time.Sleep(900 * time.Millisecond)
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
	fmt.Printf(
		"Tasarah Vineglade:\n" +
			"==================\n" +
			"The storm pushed us to the coast, we lost all control and\n" +
			"here we are. We will try to save the ship from sinking.\n" +
			"You, " + name + " and a few others are going on the land\n" +
			"and try to find a safe spot. Keep your eyes open, Harrison,\n" +
			"C'thun may be behind this heavy storm and we are still\n" +
			"waiting for the troops to arrive from the other side.\n" +
			"My crew will do everything to find where we are and how far\n" +
			"it is to the bay. Go on, we can't afford to loose time.\n")
	fmt.Scanln()
	fmt.Printf("Harrison Jones:\n" +
		"===============\n" +
		"Alright, let's go, " + name + "\n\n")
	fmt.Scanln()
}
