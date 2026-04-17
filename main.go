package main
import (
	"fmt"
	"bufio"
	"os"
	"github.com/MyLittlePico/pokedex/internal/pokecache"
	"time"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(5 * time.Second)
	conf := makeConfig()
	for{
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput( scanner.Text())
		args := make([]string, 0)
		if len(inputs) == 0 {
			continue
		}
		if len(inputs) > 1{
			args = inputs[1:]
		}
		
		command, ok :=  getCommands()[inputs[0]]
		if !ok{
			fmt.Println("Unknown command")
			continue
		}
		command.callback(&conf, &cache, args)

	}
		


}

