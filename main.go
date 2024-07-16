package main

import (
	"fmt"
	"myterraform/utility/myutility"
	"os"
)

func main() {
	fmt.Println("Starting myterraform...")

	if len(os.Args) < 3 {
		fmt.Println("Usage: myterraform <apply|plan> <config-file1> [<config-file2> ...]")
		os.Exit(1)
	}

	command := os.Args[1]
	configFiles := os.Args[2:]

	for _, configFile := range configFiles {
		fmt.Printf("Loading config from %s...\n", configFile)
		cfg, err := myutility.LoadConfig(configFile)
		if err != nil {
			fmt.Printf("\033[0;31mError loading config:\033[0m %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\033[1;33mExecuting command: %s on %s\033[0m\n", command, configFile)
		switch command {
		case "apply":
			myutility.Apply(cfg)
		case "plan":
			myutility.Plan(cfg)
		default:
			fmt.Println("Expected 'apply' or 'plan' subcommands")
			os.Exit(1)
		}
	}
}
