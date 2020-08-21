package levels

import "fmt"

// Silithuscoast level 2
func Silithuscoast() {

	SilithuscoastIntro()
}

// SilithuscoastIntro do
func SilithuscoastIntro() {
	fmt.Printf("=================================================\n" +
		"Location: Silithus coast, southwards of Ahn'Qiraj\n" +
		"=================================================\n\n\n" +
		"*You and Harrison walk from the ship to the stormy coast.\n" +
		"The waves crash onto the rocky hangs of Silithus. Your team\n" +
		"is already there trying to make a camp.*\n")
	fmt.Scanln()
	fmt.Printf("*You look up to the mountains of Silithus. The whole coast\n" +
		"of this lands are mountains with a big desert behind these\n" +
		"walls on the ocean. The thunderclouds scratch on the top*\n")
	fmt.Scanln()
	fmt.Printf("")
}
