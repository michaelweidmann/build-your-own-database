package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ExitSuccessCode = 0

type StatementType int

const (
	Insert StatementType = iota
	Select
)

type Statement struct {
	statementType StatementType
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()
		input := readInput(scanner)

		if input[0] == '.' {
			err := executeMetaCommand(input)

			if err != nil {
				fmt.Printf(err.Error())
				continue
			}
		}

		statement, err := prepareStatement(input)

		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		executeStatement(statement)
		fmt.Println("Executed.")
	}
}

func printPrompt() {
	fmt.Print("db > ")
}

func readInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()
	return input
}

func executeMetaCommand(input string) error {
	if input == ".exit" {
		os.Exit(ExitSuccessCode)
	}

	return fmt.Errorf("Unrecognized command '%s'.\n", input)
}

func prepareStatement(input string) (Statement, error) {
	if strings.HasPrefix(input, "insert") {
		return Statement{Insert}, nil
	}

	if strings.HasPrefix(input, "select") {
		return Statement{Select}, nil
	}

	return Statement{}, fmt.Errorf("Unrecognized keyword at start of '%s'.\n", input)
}
func executeStatement(statement Statement) {
	switch statement.statementType {
	case Insert:
		fmt.Println("This is an insert!")
	case Select:
		fmt.Println("This is a select!")
	}
}
