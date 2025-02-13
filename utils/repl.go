package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	Pokeapiclient    Client
	prevLocationsURL *string
	nextLocationsURL *string
}

// used to clean strings of clutter
func CleanInput(text string) []string {
	clean_text := strings.TrimSpace(text)
	clean_text = strings.ToLower(clean_text)
	clean_text_arr := strings.Fields(clean_text)

	return clean_text_arr
}

func BeginRepl(c *Config) {

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
			err := command.callback(c)
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
		"map": {
			name:        "map",
			description: "Displays the next 20 locations on the map",
			callback:    mapCallback,
		},
		"mapb": {
			name:        "map",
			description: "Displays the previous 20 locations on the map",
			callback:    mapbCallback,
		},
	}
}

// exit the program with error code 0
func exitCallback(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display all currently available commands
func helpCallback(c *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Valid commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

// next 20 locations - map command
func mapCallback(c *Config) error {
	locations_response, err := c.Pokeapiclient.getMap(c.nextLocationsURL)
	if err != nil {
		return err
	}

	// updating the two URL string stores in the config struct
	c.nextLocationsURL = locations_response.Next
	c.prevLocationsURL = locations_response.Previous

	// iterate over each struct stored in the results array of the locations struct
	// and print the name of the location in the pokemon world
	for _, loc := range locations_response.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// previous 20 locations - mapb command
func mapbCallback(c *Config) error {
	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations_response, err := c.Pokeapiclient.getMap(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locations_response.Next
	c.prevLocationsURL = locations_response.Previous

	for _, loc := range locations_response.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
