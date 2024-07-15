package main

import (
	"fmt"
	"io/ioutil"

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
