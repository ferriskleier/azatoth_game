package lib

import (
	"fmt"
)

var name string
var race string
var home string
var class string

// Player creation
func Player() {

	// Select name for player
	fmt.Printf("\nYour Name:\n>> ")
	fmt.Scan(&name)
	fmt.Printf("==============\n\n")

	// Select race for player
	Race()

	// Select class for player
	fmt.Printf("What class?\n- Warrior\n- Mage\n- Paladin\n- or type 'Help' to get more Information\n>> ")
	fmt.Scan(&class)
	if class == "Help" {
		fmt.Printf("==============\n\n")
		fmt.Printf("- Warrior:\n    A Warrior can handle heavy swords,\n    shields and fights with bloodthirst\n    until the enemies are all down\n" +
			"- Mage:\n    Mages use mana as their ressource\n    and can freeze or transform their\n    enemies directly\n" +
			"- Paladin:\n    Paladins connect to their power of light\n    to heal themselve or fight their enemies\n    with strokes of energy\n" +
			"===========================================\nWhat class?\n- Warrior\n- Mage\n- Paladin\n>> ")
		fmt.Scan(&class)
	}
	fmt.Printf("==============\n\n")

	// Summary for player
	fmt.Printf("Welcome, " + name + "\nYou left your home " + home + " to go with an expedition in\nSilithus. They are on something with old gods. C'thun,\nthe old god of madness and chaos got active and caused\nmany heroes to suffer from his strength. It's on you to\nfind his prison deep in Ahn'Qiraj, in the desert of\nSilithus. Don't fall for his tricks and try to keep a\nclear mind, he will do everything to break it.\n\n")
	fmt.Printf("==================================\n" +
		"Press Enter to start the adventure\n" +
		"==================================\n\n")
	fmt.Scanln()
}

// Race for player
func Race() {
	fmt.Printf("What race?\n- Human\n- Troll\n- Blood Elf\n- Draenei\n>> ")
	fmt.Scan(&race)
	fmt.Printf("==============\n\n")
	if race == "Human" {
		home = "Stormwind"
	} else if race == "Troll" {
		home = "Orgrimmar"
	} else if race == "Blood Elf" {
		home = "Silvermoon"
	} else if race == "Draenei" {
		home = "the Exodar"
	} else {
		fmt.Println("Invalid input, try again")
		Race()
	}
}

// Strings to export player data
func Strings() (string, string, string, string) {
	n := name
	r := race
	h := home
	c := class
	return n, r, h, c
}
