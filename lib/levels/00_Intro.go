package levels

import (
	"fmt"
)

var ask string

// Intro to ask if Intro
func Intro() {

	for i := 0; i < 100; i++ {
		fmt.Println()
	}
	fmt.Printf("Welcome to the game, do you want to play the intro\n" +
		"to learn how to play the game?\n\n")
	IntroAsk()
}

// IntroAsk to ask
func IntroAsk() {

	fmt.Printf("[Yes | No]\n>> ")
	fmt.Scan(&ask)
	fmt.Printf("===========\n\n")

	switch ask {
	case "Yes":
		IntroYes()
	case "No":
		IntroNo()
	default:
		fmt.Printf("Invalid input, try again\n")
		IntroAsk()
	}
}

// IntroYes to game
func IntroYes() {

	fmt.Printf("In this Intro I'll explain how to play the game\n\n" +
		"If you are in a part where story is told, press\n" +
		"enter to continue the dialogue. This information\n" +
		"is not given every time, so keep in mind, that if\n" +
		"the dialogue is paused, you can press enter, to\n  ...")
	fmt.Scanln()
	fmt.Printf("  ...\n\ncontinue the game, very good.\n")
	fmt.Scanln()
	fmt.Printf("When the game gives you a choice, just type one of\n" +
		"the given answers like you did when I asked you if\n" +
		"you want to play this intro.\n" +
		"This is a simple Quest, you have to go to the next\n" +
		"step, okay?\n\n")
	IntroQuest()
	fmt.Printf("Very Good, let's talk about fighting\n")
	fmt.Scanln()
	fmt.Printf("You will now see a simple fight. The parts of the\n" +
		"fight screen are:\n")
	fmt.Scanln()
	fmt.Printf("Your healthbar:\n\n" +
		"[[Your Name]]\n" +
		"[# # # # # # # # # #]\n(you've 10HP in this example, 10x '#'\n")
	fmt.Scanln()
	fmt.Printf("Your enemy's healthbar:\n\n" +
		"[[Enemy Name]]\n" +
		"[@ @ @]\n(the enemy has 3HP in this example, 3x '@'\n")
	fmt.Scanln()
	fmt.Printf("Here you get all your possible attacks, as with questions\n" +
		"you can type the number of what attack you want to use\n" +
		"Here are your attacks:\n\n")
	FightQuest()
	fmt.Printf("\nThe fight is over\n")
	fmt.Scanln()
	fmt.Printf("You will gain more and more attacks over time\n" +
		"Your attacks depend on your class and Items\n" +
		"Items you could use are displayed with your attacks\n")
	fmt.Scanln()
	fmt.Printf("Now you have one more thing to do: start the game\n" +
		"To do this, answer this type of question you will see\n" +
		"often in this game\n\n" +
		"What do you want to do?\n\n")
	IntroEnd()
}

// IntroQuest do
func IntroQuest() {
	fmt.Printf("[Accept | No]\n>> ")
	var quest1 string
	fmt.Scan(&quest1)
	switch quest1 {
	case "Accept":
		fmt.Printf("\n")
	case "No":
		fmt.Printf("\nWell, you have no other choice, let's try it\n" +
			"again:\n\n")
		IntroQuest()
	default:
		fmt.Printf("\nInvalid input, try again\n\n")
		IntroQuest()
	}
}

// IntroNo to game
func IntroNo() {

	fmt.Printf("Alright, press enter to begin\n")
	fmt.Scanln()
}

// FightQuest do
func FightQuest() {
	var choiceFight string
	fmt.Printf("[1] Sword attack\n" +
		"[2] Flee\n>> ")
	fmt.Scan(&choiceFight)
	switch choiceFight {
	case "1":
		fmt.Printf("\nYou attack the enemy with your sword\n" +
			"Enemy loses 3HP\n")
		fmt.Scanln()
		fmt.Printf("The enemy is defeated. Congratulations!\n")
	case "2":
		fmt.Printf("\nYou escaped the fight, not very brave\n")
	default:
		fmt.Printf("\nInvalid input, try again\n\n")
		FightQuest()
	}
}

// IntroEnd ready
func IntroEnd() {
	var end string
	fmt.Printf("[Start the game | Replay Intro]\n>> ")
	fmt.Scan(&end)
	switch end {
	case "Start":
		for i := 0; i < 8; i++ {
			fmt.Scanln()
		}
	case "Replay":
		for i := 0; i < 5; i++ {
			fmt.Scanln()
		}
		Intro()
	default:
		fmt.Printf("\nInvalid input, try again\n\n")
		IntroEnd()
	}
}
