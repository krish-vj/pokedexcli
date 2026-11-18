package main

import (
	"fmt"

	"github.com/krish-vj/pokedexcli/resps"
)
func commandExplore(cfg *config, name ...string) error{
	if (len(name)==0){
		return fmt.Errorf("pass the area name or id bozo\n")
	}
	fmt.Printf("Exploring %s...\n", name[0])
	url:="https://pokeapi.co/api/v2/location-area/"+name[0]
	pokis, err:=lookUp[resps.PokemonsInArea](url, cfg)
	if (err!=nil){
		fmt.Print("kay zala hey?", err)
		return err;
	}
	fmt.Println("Found Pokemon:")
	for _, idk:= range pokis.PokemonEncounters{
		fmt.Printf(" - %s\n", idk.Pokemon.Name)
	}
	return nil;
}