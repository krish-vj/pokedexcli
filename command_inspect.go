package main

import "fmt"

func commandInspect(cfg *config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("Pass the name of pokemon\n")

	}
	pokemon, ok:=cfg.caughtPokemons[name[0]]
	if (!ok){
		fmt.Println("you have not caught that pokemon")
	}else{
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d", pokemon.Name,  pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stats:=range pokemon.Stats{
			fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
		}
		fmt.Println("Types:")
		for _, ptype:=range pokemon.Types{
			fmt.Printf("  - %s\n", ptype.Type.Name)
		}
	}	
	return nil;
}