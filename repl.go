package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/krish-vj/pokedexcli/pokecache"
	"github.com/krish-vj/pokedexcli/resps"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct{
	Next string 
	Previous string 
	cachePtr *pokecache.Cache
	caughtPokemons map[string]resps.Pokemon
}
func getCofig() *config{
	return &config{caughtPokemons: make(map[string]resps.Pokemon)};
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
		"explore":{
			name: "explore",
			description: "Explore an area to see pokemon(s) in that area",
			callback: commandExplore,
		},
		"catch":{
			name: "catch",
			description: "Catch a pokemon by passing it's name/id",
			callback: commandCatch,
		},
		"inspect":{
			name: "inspect",
			description: "Inspect a pokemon by passing it's name/id",
			callback: commandInspect,
		},
		"pokedex":{
			name: "pokedex",
			description: "prints a list of all the names of the Pokemon you have caught",
			callback: commandPokedex,
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
			var argsToPass []string
			if (len(res)>1){
				argsToPass=res[1:]
			}
			cmdStruct, ok:=getCommands()[gotCmd]
			if (!ok){
				fmt.Print("Unknown command\n");
				continue;
			}else{
				err:=cmdStruct.callback(cfg, argsToPass...)
				if (err!=nil){
					fmt.Print(err)
				}
			}
		}

	}
}
