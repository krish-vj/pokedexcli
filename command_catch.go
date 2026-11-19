package main

import (
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/krish-vj/pokedexcli/resps"
)
func getProb(baseExp int) float64{
	const K=350
	return K/float64(K+baseExp)
}

func commandCatch(cfg *config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("pass a pokemon name or id")
	}
	url:="https://pokeapi.co/api/v2/pokemon/"+name[0]
	
	pokemon, err:=lookUp[resps.Pokemon](url, cfg)
	if (err!=nil){
		return err;
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name[0])
	probCatching:=getProb(pokemon.BaseExperience)
	slog.Info("Attempting to catch pokemon", "name", name[0], "Chance", probCatching)
	if (rand.Float64()<probCatching){
		cfg.caughtPokemons[name[0]]=*pokemon
		fmt.Printf("%s was caught!\n", name[0]);
		fmt.Println("You may now inspect it with the inspect command.")
	}else{
		fmt.Printf("%s escaped\n!", name[0]);
	}
	return nil;
}