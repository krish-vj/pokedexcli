package main

import (
	"fmt"
	"os"
)
func commandExit(cfg *config, _ ...string) error{
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("how has the program not terminated yet?");
}
