package main

import (
	"fmt"
	"os"
	"proton/core"
)

func main() {
	config, err := core.GetProtonConfig("proton.toml")
	if err != nil {
		fmt.Printf("error while loading config: %v\n", err.Error())
		os.Exit(1)
	}
	core.Setup(&config)
}
