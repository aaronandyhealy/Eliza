package main

import (
	"fmt"
	"eliza"
	
)


func main() {
	fmt.Println("How do you do. Please tell me what's on your mind.")
	defer fmt.Println("Goodbye. It was nice talking to you.")

	fmt.Println()

	for {
		input := eliza.GetInput()

		parsed := eliza.ParseInput(input)

		if quit := eliza.CheckForQuit(parsed); quit {
			break
		}

		fmt.Println(parsed)
	}
}