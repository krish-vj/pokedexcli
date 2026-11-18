package main

import (
	"fmt"
	"github.com/krish-vj/pokedexcli/resps"
	
)

func commandMap(cfg *config, _ ...string) error{
	url:="https://pokeapi.co/api/v2/location-area?offset=0&limit=20";
	if (cfg.Next!=""){
		url=cfg.Next;
	}
	pokeStr, err:=lookUp[resps.PokeApiResp](url, cfg)
	if (err!=nil || pokeStr==nil){
		return err
	}
	cfg.Next=pokeStr.Next;
	if (pokeStr.Previous!=nil ){
		cfg.Previous=pokeStr.Previous.(string)
	}
	// now print the current 20
	for _, res:= range pokeStr.Results{
		fmt.Println(res.Name)
	}

	return nil;
}