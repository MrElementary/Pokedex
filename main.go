package main

import (
	"time"

	"github.com/MrElementary/Pokedex/utils"
)

// we're using 5 minutes as the interval period to clean the cache
// the 5 seconds is for the http.Client object to time out our http connection
func main() {
	pokeClient := utils.NewClient(5*time.Second, time.Minute*5)
	c := &utils.Config{
		Pokeapiclient: pokeClient,
	}
	utils.BeginRepl(c)
}
