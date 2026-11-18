package main

import (
	"fmt"
)
func commandMapB(cfg *config) error{
	if (cfg.Previous==""){
		fmt.Print("you're on the first page");
		return nil;
	}	
		// should i call the api again? or should i store the results locally in a slice of 
		// Results? //ig i will stick with calling the api for now

	url:= cfg.Previous;
	pokeStr, err:=lookUp(url, cfg)
	if (err!=nil){
		return err;
	}
	cfg.Next=pokeStr.Next;
	
	if (pokeStr.Previous!=nil ){
		cfg.Previous=pokeStr.Previous.(string)
	}else{
		cfg.Previous=""
	}
	// now print the current 20
	for _, res:= range pokeStr.Results{
		fmt.Println(res.Name)
	}

	return nil;


}