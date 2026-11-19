package main

import "fmt"

func commandPokedex(cfg *config, _ ...string)error {
	if (len(cfg.caughtPokemons)==0){
		fmt.Println("You haven't caught any pokemons yet. Use 'catch <pokemon_name>' command to catch")
		return nil;
	}
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.caughtPokemons{
		fmt.Printf("  - %s\n", key)
	}
	return nil
}