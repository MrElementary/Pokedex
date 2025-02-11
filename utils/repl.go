package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// used to clean strings of clutter
func CleanInput(text string) []string {
	clean_text := strings.TrimSpace(text)
	clean_text = strings.ToLower(clean_text)
	clean_text_arr := strings.Fields(clean_text)

	return clean_text_arr
}

func BeginRepl() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type 'exit' to quit")

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()
		input_arr := CleanInput(input)
		command_key := input_arr[0]

		command, ok := getCommands()[command_key]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Print("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCallback,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCallback,
		},
	}
}

func exitCallback() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCallback() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Valid commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
