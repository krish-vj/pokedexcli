package main

import (
	"github.com/krish-vj/pokedexcli/pokecache"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct{
	Next string 
	Previous string 
	cachePtr *pokecache.Cache
}
func getCofig() *config{
	return &config{};
}

type PokeApiResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}



func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map":{
			name: "map",
			description: "Fetches 20 locations",
			callback: commandMap,
		},
		"mapb":{
			name: "mapb",
			description: "Fetches last 20",
			callback: commandMapB,
		},
	}
}
func cleanInput(text string) []string{
	
	var res []string 
	var b strings.Builder
	for _, char:=range text{
		if (char==' '){
			if (b.Len()!=0){
				res=append(res, b.String())
				b.Reset()
			}
		}else{
			b.WriteRune(char)
		}
		
	}
	if (b.Len()!=0){
		res=append(res, b.String())
	}

	return res
}
func startRepl(cfg *config){

	// fmt.Println("starting scanner...")
	scanner:=bufio.NewScanner(os.Stdin)
	for true{
		fmt.Print("Pokedex > ")
		// scanner.Scan()
		//blocking call; waits till we encounter \n or it's equivalents 
		if !scanner.Scan(){
			break
		}
		res:=cleanInput(scanner.Text())
		if len(res)==0{
			fmt.Println("got nothing bro :(")
		}else{
			gotCmd:=strings.ToLower(res[0]);
			cmdStruct, ok:=getCommands()[gotCmd]
			if (!ok){
				fmt.Print("Unknown command\n");
				continue;
			}else{
				err:=cmdStruct.callback(cfg)
				if (err!=nil){
					fmt.Print(err)
				}
			}
		}

	}
}
