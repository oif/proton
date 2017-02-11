package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"proton/core"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	config, err := core.GetProtonConfig("proton.toml")
	if err != nil {
		fmt.Printf("error while loading config: %v\n", err.Error())
		os.Exit(1)
	}
	core.Setup(&config)
}
