package main

import(
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"io"
	"github.com/MyLittlePico/pokedex/internal/pokeAPI"
)

type cliCommand struct{
	name string
	description string
	callback func(conf *config) error
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


func commandExit(conf *config) error{
	fmt.Println("Closing the Pokedex... Goodbye!")	
	os.Exit(0)
	return nil
}

func commandHelp(conf *config) error{
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range getCommands(){
		fmt.Printf("%s: %s\n",cmd.name, cmd.description)
	}
	return nil
}

func commandMap(conf *config) error{
	url := conf.nextUrl
	res, err := http.Get(url)
	if err != nil{
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	
	var data pokeapi.LocationAreas
	err = json.Unmarshal(body, &data)
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
func commandMapb(conf *config) error{
	url := conf.previousUrl
	res, err := http.Get(url)
	if err != nil{
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	
	var data pokeapi.LocationAreas
	err = json.Unmarshal(body, &data)
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

