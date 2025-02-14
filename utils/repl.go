package utils

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	Pokeapiclient    Client
	prevLocationsURL *string
	nextLocationsURL *string
	CaughtPokemon    map[string]Pokemon
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

		args := []string{}
		if len(input_arr) > 1 {
			args = input_arr[1:]
		}

		command, ok := getCommands()[command_key]
		if ok {
			err := command.callback(c, args...)
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
		"explore": {
			name:        "explore <location_name>",
			description: "Explore the pokemon at the given location name",
			callback:    exploreCallback,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Make an attempt to try and catch the pokemon",
			callback:    catchCallback,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "shows details about a caught pokemon",
			callback:    inspectCallback,
		},
		"pokedex": {
			name:        "pokedex",
			description: "shows all pokemon that's been caught",
			callback:    pokedexCallback,
		},
	}
}

// exit the program with error code 0
func exitCallback(c *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// display all currently available commands
func helpCallback(c *Config, args ...string) error {
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
func mapCallback(c *Config, args ...string) error {
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
func mapbCallback(c *Config, args ...string) error {
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

// get all pokemon at the specific location name - explore command
func exploreCallback(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := c.Pokeapiclient.getSingleLocation(name)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}

func catchCallback(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := c.Pokeapiclient.getPokemon(name)

	if err != nil {
		return err
	}

	base_exp := pokemon.BaseExperience

	exp_to_chance := int(float64(base_exp) / 255 * 100)
	player_roll := rand.Intn(100)

	fmt.Printf("Throwing a Pokeball at %s...\nYou need to roll above %d\nYou rolled %d\n", pokemon.Name, exp_to_chance, player_roll)

	if player_roll <= exp_to_chance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)

	c.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}

func inspectCallback(c *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := c.CaughtPokemon[name]

	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}

	return nil
}

func pokedexCallback(c *Config, args ...string) error {
	if len(c.CaughtPokemon) == 0 {
		fmt.Printf("You haven't caught anything yet\n")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for name := range c.CaughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
