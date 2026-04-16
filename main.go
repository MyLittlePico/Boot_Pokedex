package main
import (
	"fmt"
	"bufio"
	"os"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput( scanner.Text())
		if len(inputs) == 0 {
			continue
		}
		command, ok := cliCommands[inputs[0]]
		if !ok{
			fmt.Println("Unknown command")
			continue
		}
		command.callback()

	}
		


}

