package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"

	"github.com/MyLittlePico/pokedex/internal/pokeAPI"
	"github.com/MyLittlePico/pokedex/internal/pokecache"
)

type cliCommand struct{
	name string
	description string
	callback func(conf *config, cache *pokecache.Cache, args []string) error
}
type config struct{
	nextUrl string
	previousUrl string
	pokedex  map[string]pokeapi.PokemonStat
}

func makeDex() map[string]pokeapi.PokemonStat{
	return make(map[string]pokeapi.PokemonStat)
}

func makeConfig () config {
	c := config{
		previousUrl: "https://pokeapi.co/api/v2/location-area",
		nextUrl: "https://pokeapi.co/api/v2/location-area",
		pokedex: makeDex(),
	}
	return c
} 
 
func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit":{
			name: "exit",
			description: "Exit the Pokedex",
			callback:	commandExit,
		},
		"help":{
			name: "help",
			description: "Displays a help message",
			callback:	commandHelp,
		},
		"map":{
			name: "map",
			description: "Displays next 20 areas",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Displays previous 20 areas",
			callback: commandMapb,
		},
		"explore":{
			name: "explore",
			description: "Displays Pokémons locate in an area",
			callback: commandExplore,
		},
		"catch":{
			name :"catch",
			description: "Catch Pokemon and add to the Pokedex.",
			callback: commandCatch,
		},
		"inspect":{
			name :"inspect",
			description:"Display details about a already caught Pokemon",
			callback: commandInspect,
		},
		"pokedex":{
			name :"pokedex",
			description: "go to hell",
		},
	}

}


func commandExit(conf *config, cache *pokecache.Cache, args []string) error{
	fmt.Println("Closing the Pokedex... Goodbye!")	
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, cache *pokecache.Cache, args []string) error{
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n",cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *config, cache *pokecache.Cache, args []string) error{
	url := conf.nextUrl
	res, err := cache.Url(url)
	if err != nil{
		return err
	}
	return processLocationAreasData(conf, res)
}

func commandMapb(conf *config, cache *pokecache.Cache, args []string) error{
	url := conf.previousUrl
	res, err := cache.Url(url)
	if err != nil{
		return err
	}
	return processLocationAreasData(conf, res)
}
func processLocationAreasData (conf *config, res []byte) error{
	var data pokeapi.LocationAreas
	err := json.Unmarshal(res, &data)
	if err != nil {
		return err
	}
	conf.previousUrl = data.Previous
	conf.nextUrl = data.Next
	for _, locations := range data.Results {
		fmt.Println(locations.Name)
	} 
	return nil
}

func commandExplore(conf *config, cache *pokecache.Cache, args []string) error{
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", args[0])
	res, err := cache.Url(url)
	if err != nil {
		return err
	}
	return processAreaDetailData(res)

}

func processAreaDetailData(res []byte) error{
	var data pokeapi.AreaDetail
	err := json.Unmarshal(res, &data)
	if err != nil {
		return err
	}
	for _, info := range data.PokemonEncounters{
		fmt.Printf("-%s\n", info.Pokemon.Name)
	}
	return nil
}

func commandCatch(conf *config, cache *pokecache.Cache, args []string) error{
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", name)
	res, err := cache.Url(url)
	if err != nil {
		return err
	}
	err = processPokemonData(conf ,res)
	if err != nil {
		return err
	}
	
	return nil
}

func processPokemonData(conf *config, res []byte) error{
	var data pokeapi.PokemonStat
	err := json.Unmarshal(res, &data)
	if err != nil {
		return  err
	}

	if  data.BaseExperience > rand.Intn(300) {
		fmt.Printf("%s escaped!\n", data.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", data.Name)
	conf.pokedex[data.Name] = data
	return nil
}

func commandInspect(conf *config, cache *pokecache.Cache, args []string) error{
	name := args[0]
	pokemon ,ok := conf.pokedex[name]
	if !ok{
		fmt.Println("you have not caught that pokemon")
	}
	fmt.Printf("Name: %v\n",pokemon.Name)
	fmt.Printf("Height: %v\n",pokemon.Height)
	fmt.Printf("Weight: %v\n",pokemon.Weight)
	fmt.Println("Stats:")
	for _,stat := range pokemon.Stats{
		fmt.Printf("\t-%v: %v\n",stat.Stat.Name,stat.BaseStat)
	}
	fmt.Println("Types:")
	for _,theType := range pokemon.Types{
		fmt.Printf("\t-%v\n",theType.Type.Name)
	}
	return nil
}