package main

import (
	"fmt"
	"github.com/oif/proton/core"
	"os"
)

func main() {
	config, err := core.GetProtonConfig("proton.toml")
	if err != nil {
		fmt.Printf("error while loading config: %v\n", err.Error())
		os.Exit(1)
	}
	core.Setup(&config)
}
