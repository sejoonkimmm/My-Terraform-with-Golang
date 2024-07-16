package myutility

import (
	"fmt"
	"myterraform/modules/mymodule"
	"myterraform/providers/myprovider"
)

func Apply(cfg *Config) {
	fmt.Println("\033[1;32mApplying infrastructure changes...\033[0m")
	provider := myprovider.NewMyProvider(cfg.Provider.Region)

	for _, resource := range cfg.Resources {
		state := mymodule.NewStateManager(resource.Name + ".tfstate")
		state.LoadState()

		if state.ResourceExists(resource.Name) {
			fmt.Printf("\033[0;32mResource %s already exists, skipping...\033[0m\n", resource.Name)
			continue
		}

		fmt.Printf("Resource Type: %s, Name: %s\n", resource.Type, resource.Name)

		if resource.Type == "aws_instance" {
			fmt.Printf("\033[1;34mCreating AWS EC2 instance: AMI=%s, InstanceType=%s\033[0m\n", resource.Ami, resource.InstanceType)
			module := mymodule.NewMyModule(provider, resource.Ami, resource.InstanceType)
			module.Apply()
			state.AddResource(resource.Name, mymodule.ResourceState{
				Type:         resource.Type,
				Name:         resource.Name,
				Ami:          resource.Ami,
				InstanceType: resource.InstanceType,
			})
		}
		state.SaveState()
	}
	fmt.Println("\033[0;32mApply complete!\033[0m")
}

func Plan(cfg *Config) {
	fmt.Println("\033[1;33mPlanning infrastructure changes...\033[0m")
	provider := myprovider.NewMyProvider(cfg.Provider.Region)

	for _, resource := range cfg.Resources {
		state := mymodule.NewStateManager(resource.Name + ".tfstate")
		state.LoadState()

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
