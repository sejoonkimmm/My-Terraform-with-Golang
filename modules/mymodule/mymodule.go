package mymodule

import (
	"fmt"
	"myterraform/providers/myprovider"
)

type MyModule struct {
	provider     *myprovider.MyProvider
	Ami          string
	InstanceType string
}

func NewMyModule(provider *myprovider.MyProvider, ami string, instanceType string) *MyModule {
	fmt.Println("Creating new MyModule instance...")
	return &MyModule{provider: provider, Ami: ami, InstanceType: instanceType}
}

func (m *MyModule) Apply() {
	fmt.Println("Applying MyModule")
	_, err := m.provider.CreateInstance(m.Ami, m.InstanceType)
	if err != nil {
		fmt.Printf("Error applying module: %v\n", err)
	}
}

func (m *MyModule) Plan() {
	fmt.Println("Planning MyModule")
	m.provider.PlanInstance(m.Ami, m.InstanceType)
}
