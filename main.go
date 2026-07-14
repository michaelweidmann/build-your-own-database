package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		printPrompt()
		input := readInput()

		if input == ".exit" {
			return
		}

		fmt.Printf("Unrecognized command '%s'.\n", input)
	}
}

func printPrompt() {
	fmt.Print("db > ")
}

func readInput() string {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatalln("Unexpected error occurred: ", err)
	}

	return input
}
