package levels

import "fmt"

// Intro to explain game
func Intro() {
	var ask string
	for i := 0; i < 100; i++ {
		fmt.Println()
	}
	fmt.Printf("Welcome to the game, do you want to play the intro\n" +
		"to learn how to play the game?\n\n" +
		"[Yes or No]\n>> ")
	fmt.Scan(&ask)
	fmt.Println("===========")
}
