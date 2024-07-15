package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting myterraform...")

	if len(os.Args) < 3 {
		fmt.Println("Usage: myterraform <apply|plan> <config-file>")
		os.Exit(1)
	}

	command := os.Args[1]
	configFile := os.Args[2]

	fmt.Printf("Loading config from %s...\n", configFile)
	config, err := loadConfig(configFile)
	if err != nil {
		fmt.Printf("\033[0;31mError loading config:\033[0m %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\033[1;33mExecuting command: %s\033[0m\n", command)
	switch command {
	case "apply":
		apply(config)
	case "plan":
		plan(config)
	default:
		fmt.Println("Expected 'apply' or 'plan' subcommands")
		os.Exit(1)
	}
}
