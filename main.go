package main

import (
	"github.com/krish-vj/pokedexcli/pokecache"
	"fmt"
	"log/slog"
	"time"
)

func main(){
	logF, err:=initLogger("new.log")
	if (err!=nil){
		fmt.Print("Failed to open log file", err);
	}else{
		defer logF.Close()
	}
	const timeDurationForCache=time.Minute*5;
	slog.Info("creating config")
	configPtr:=getCofig()
	cachePtr:=pokecache.NewCache(timeDurationForCache);
	configPtr.cachePtr=cachePtr
	startRepl(configPtr)
}
