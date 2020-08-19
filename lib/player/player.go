package player

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
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
	fmt.Printf("=================================================\n" +
		"Wrath of C'Thun - World of Warcaft Text Adventure\n" +
		"=================================================\n\n\n" +
		"\nYour Name:\n>> ")
	fmt.Scan(&name)
	fmt.Printf("==============\n\n")

	Race()
	Class()

	// Summary for player
	fmt.Printf("==================================\n" +
		"Press enter to start the adventure\n" +
		"==================================\n\n\n")
	fmt.Scanln()
}

// Race for player
func Race() {
	fmt.Printf("What race?\n- Human\n- Troll\n- Blood Elf\n- Draenei\n>> ")
	fmt.Scanln(&race)
	fmt.Printf("==============\n\n")
	if race == "Human" {
		home = "Stormwind"
	} else if race == "Troll" {
		home = "Orgrimmar"
	} else if race == "Blood" {
		fmt.Scanln()
		fmt.Scanln()
		race = "Blood Elf"
		home = "Silvermoon"
	} else if race == "Draenei" {
		home = "the Exodar"
	} else {
		fmt.Println("Invalid input, try again\n" + "==============\n\n")
		Race()
	}
}

// Class for player
func Class() {
	var newclass string
	fmt.Printf("What class?\n- Warrior\n- Hunter\n- Mage\n- Paladin\n- or type 'Help' to get more Information\n>> ")
	fmt.Scan(&newclass)
	fmt.Printf("==============\n\n")
	if newclass == "Warrior" {
		class = "Warrior"
	} else if newclass == "Hunter" {
		class = "Hunter"
	} else if newclass == "Mage" {
		class = "Mage"
	} else if newclass == "Paladin" {
		class = "Paladin"
	} else if newclass == "Help" {
		fmt.Printf("- Warrior:\n    A Warrior can handle heavy swords,\n    shields and fights with bloodthirst\n    until the enemies are all down\n" +
			"- Hunter:\n    Hunters tame wild animals to fight\n    side by side with them. Your animal\n    depends on your chosen race\n" +
			"- Mage:\n    Mages use mana as their ressource\n    and can freeze or transform their\n    enemies directly\n" +
			"- Paladin:\n    Paladins connect to their power of light\n    to heal themselve or fight their enemies\n    with strokes of energy\n" +
			"===========================================\n\n")
		Class()
	} else {
		fmt.Println("Invalid input, try again\n" + "==============\n\n")
		Class()
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
