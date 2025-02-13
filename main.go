package main

import (
	"time"

	"github.com/MrElementary/Pokedex/utils"
)

func main() {
	pokeClient := utils.NewClient(5 * time.Second)
	c := &utils.Config{
		Pokeapiclient: pokeClient,
	}
	utils.BeginRepl(c)
}
