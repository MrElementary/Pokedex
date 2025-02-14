package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/MrElementary/Pokedex/internal"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://pokeapi.co/api/v2/location-area",
			val: []byte(`{"count":1054,"next":"https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			"previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},
			{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},
			{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},
			{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},
			{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},
			{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},
			{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},
			{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},
			{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},
			{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},
			{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},
			{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},
			{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},
			{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},
			{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},
			{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},
			{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},
			{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},
			{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},
			{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := internal.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	cache := internal.NewCache(baseTime)
	cache.Add("https://pokeapi.co/api/v2/location-area", []byte(`{"count":1054,"next":"https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
	"previous":null,"results":[{"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"},
	{"name":"eterna-city-area","url":"https://pokeapi.co/api/v2/location-area/2/"},
	{"name":"pastoria-city-area","url":"https://pokeapi.co/api/v2/location-area/3/"},
	{"name":"sunyshore-city-area","url":"https://pokeapi.co/api/v2/location-area/4/"},
	{"name":"sinnoh-pokemon-league-area","url":"https://pokeapi.co/api/v2/location-area/5/"},
	{"name":"oreburgh-mine-1f","url":"https://pokeapi.co/api/v2/location-area/6/"},
	{"name":"oreburgh-mine-b1f","url":"https://pokeapi.co/api/v2/location-area/7/"},
	{"name":"valley-windworks-area","url":"https://pokeapi.co/api/v2/location-area/8/"},
	{"name":"eterna-forest-area","url":"https://pokeapi.co/api/v2/location-area/9/"},
	{"name":"fuego-ironworks-area","url":"https://pokeapi.co/api/v2/location-area/10/"},
	{"name":"mt-coronet-1f-route-207","url":"https://pokeapi.co/api/v2/location-area/11/"},
	{"name":"mt-coronet-2f","url":"https://pokeapi.co/api/v2/location-area/12/"},
	{"name":"mt-coronet-3f","url":"https://pokeapi.co/api/v2/location-area/13/"},
	{"name":"mt-coronet-exterior-snowfall","url":"https://pokeapi.co/api/v2/location-area/14/"},
	{"name":"mt-coronet-exterior-blizzard","url":"https://pokeapi.co/api/v2/location-area/15/"},
	{"name":"mt-coronet-4f","url":"https://pokeapi.co/api/v2/location-area/16/"},
	{"name":"mt-coronet-4f-small-room","url":"https://pokeapi.co/api/v2/location-area/17/"},
	{"name":"mt-coronet-5f","url":"https://pokeapi.co/api/v2/location-area/18/"},
	{"name":"mt-coronet-6f","url":"https://pokeapi.co/api/v2/location-area/19/"},
	{"name":"mt-coronet-1f-from-exterior","url":"https://pokeapi.co/api/v2/location-area/20/"}]}`))

	_, ok := cache.Get("https://pokeapi.co/api/v2/location-area")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
}
