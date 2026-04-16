package main

import(
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"io"
	"github.com/MyLittlePico/pokedex/internal/pokeAPI"
	"github.com/MyLittlePico/pokedex/internal/pokecache"
)

type cliCommand struct{
	name string
	description string
	callback func(conf *config, cache *pokecache.Cache) error
}
type config struct{
	nextUrl string
	previousUrl string
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
	}

}


func commandExit(conf *config, cache *pokecache.Cache) error{
	fmt.Println("Closing the Pokedex... Goodbye!")	
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, cache *pokecache.Cache) error{
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n",cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *config, cache *pokecache.Cache) error{
	url := conf.nextUrl
	val, ok := cache.Get(url)
	var data []byte
	if ok{
		data = val
	}else {
		res, err := http.Get(url)
		if err != nil{
			return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil{
			return err
		}
		cache.Add(url, data)
	}
	var unmarshaledData pokeapi.LocationAreas
	err := json.Unmarshal(data, &unmarshaledData)
	if err != nil {
		return err
	}
	conf.previousUrl = unmarshaledData.Previous
	conf.nextUrl = unmarshaledData.Next
	for _, locations := range unmarshaledData.Results {
		fmt.Println(locations.Name)
	} 
	return nil

}

func commandMapb(conf *config, cache *pokecache.Cache) error{
	url := conf.previousUrl
	val, ok := cache.Get(url)
	var data []byte
	if ok{
		data = val
	}else {
		res, err := http.Get(url)
		if err != nil{
			return err
		}
		defer res.Body.Close()
		data, err = io.ReadAll(res.Body)
		if err != nil{
			return err
		}
		cache.Add(url, data)
	}
	var unmarshaledData pokeapi.LocationAreas
	err := json.Unmarshal(data, &unmarshaledData)
	if err != nil {
		return err
	}
	conf.previousUrl = unmarshaledData.Previous
	conf.nextUrl = unmarshaledData.Next
	for _, locations := range unmarshaledData.Results {
		fmt.Println(locations.Name)
	} 
	return nil

}
