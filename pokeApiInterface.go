package main

//baad mai implement karenge
import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func lookUp[T any](url string, cfg *config) (*T, error) {
	var pokeStr T
	//check cache
	bytes, mila := cfg.cachePtr.Get(url)
	done := false
	if mila {
		slog.Info("cache hit!", "key", url)
		err := json.Unmarshal(bytes, &pokeStr)
		//if err go into the else block
		if err == nil {
			done = true
		}
	}else{
		slog.Info("cache MISS!", "key", url)
	}
	if !done {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		// decoder := json.NewDecoder(resp.Body)
		bytesS, err:=io.ReadAll(resp.Body)
		// err = decoder.Decode(&pokeStr)
		if err != nil {
			return nil, err
		}
		// bytes, err := json.Marshal(&pokeStr)
		json.Unmarshal(bytesS, &pokeStr)
		cfg.cachePtr.Add(url, bytesS)
		
	}
	return &pokeStr, nil
}
