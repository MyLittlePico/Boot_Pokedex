package main
import (
	"fmt"
	"bufio"
	"os"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		previousUrl: "https://pokeapi.co/api/v2/location-area",
		nextUrl: "https://pokeapi.co/api/v2/location-area",
	}
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput( scanner.Text())
		if len(inputs) == 0 {
			continue
		}
		command, ok :=  getCommands()[inputs[0]]
		if !ok{
			fmt.Println("Unknown command")
			continue
		}
		command.callback(&conf)

	}
		


}

