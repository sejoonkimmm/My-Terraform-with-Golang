package main

import (
	"fmt"
	"io/ioutil"
	"myterraform/modules/mymodule"
	"myterraform/providers/myprovider"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Provider struct {
		Region string `yaml:"region"`
	} `yaml:"provider"`
	Resources []Resource `yaml:"resources"`
}

type Resource struct {
	Type         string `yaml:"type"`
	Name         string `yaml:"name"`
	Ami          string `yaml:"ami"`
	InstanceType string `yaml:"instance_type"`
}

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
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	fmt.Printf("Executing command: %s\n", command)
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

func loadConfig(filename string) (*Config, error) {
	fmt.Printf("Reading config file: %s\n", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fmt.Println("Parsing config file...")
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	fmt.Println("Config loaded successfully")
	return &config, nil
}

func apply(config *Config) {
	fmt.Println("Applying infrastructure changes...")
	provider := myprovider.NewMyProvider(config.Provider.Region)
	for _, resource := range config.Resources {
		fmt.Printf("Resource Type: %s, Name: %s\n", resource.Type, resource.Name) // 디버깅용 출력

		if resource.Type == "aws_instance" {
			fmt.Printf("Creating AWS EC2 instance: AMI=%s, InstanceType=%s\n", resource.Ami, resource.InstanceType)
			module := mymodule.NewMyModule(provider, resource.Ami, resource.InstanceType)
			module.Apply()
		}
	}
}

func plan(config *Config) {
	fmt.Println("Planning infrastructure changes...")
	provider := myprovider.NewMyProvider(config.Provider.Region)
	for _, resource := range config.Resources {
		fmt.Printf("Resource Type: %s, Name: %s\n", resource.Type, resource.Name) // 디버깅용 출력

		if resource.Type == "aws_instance" {
			fmt.Printf("Planning to create AWS EC2 instance: AMI=%s, InstanceType=%s\n", resource.Ami, resource.InstanceType)
			module := mymodule.NewMyModule(provider, resource.Ami, resource.InstanceType)
			module.Plan()
		}
	}
}
