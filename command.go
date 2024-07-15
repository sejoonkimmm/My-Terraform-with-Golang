package main

import (
	"fmt"
	"myterraform/modules/mymodule"
	"myterraform/providers/myprovider"
)

func apply(config *Config) {
	fmt.Println("Applying infrastructure changes...")
	provider := myprovider.NewMyProvider(config.Provider.Region)
	state := mymodule.NewStateManager()
	state.LoadState()

	for _, resource := range config.Resources {
		if state.ResourceExists(resource.Name) {
			fmt.Printf("\033[0;32mResource %s already exists, skipping...\033[0m\n", resource.Name)
			continue
		}

		fmt.Printf("Resource Type: %s, Name: %s\n", resource.Type, resource.Name)

		if resource.Type == "aws_instance" {
			fmt.Printf("\033[1;34mCreating AWS EC2 instance: AMI=%s, InstanceType=%s\033[0m\n", resource.Ami, resource.InstanceType)
			module := mymodule.NewMyModule(provider, resource.Ami, resource.InstanceType)
			module.Apply()
			state.AddResource(resource.Name)
		}
	}
	state.SaveState()
	fmt.Println("\033[0;32mApply complete!\033[0m")
}

func plan(config *Config) {
	fmt.Println("Planning infrastructure changes...")
	provider := myprovider.NewMyProvider(config.Provider.Region)
	state := mymodule.NewStateManager()
	state.LoadState()

	for _, resource := range config.Resources {
		if state.ResourceExists(resource.Name) {
			fmt.Printf("\033[0;32mResource %s already exists, skipping plan...\033[0m\n", resource.Name)
			continue
		}

		fmt.Printf("Resource Type: %s, Name: %s\n", resource.Type, resource.Name)

		if resource.Type == "aws_instance" {
			fmt.Printf("\033[1;34mPlanning to create AWS EC2 instance: AMI=%s, InstanceType=%s\033[0m\n", resource.Ami, resource.InstanceType)
			module := mymodule.NewMyModule(provider, resource.Ami, resource.InstanceType)
			module.Plan()
		}
	}
	fmt.Println("\033[0;32mPlan complete!\033[0m")
}
