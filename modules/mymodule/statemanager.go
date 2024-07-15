package mymodule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StateManager struct {
	stateFile string
	State     map[string]bool
}

func NewStateManager() *StateManager {
	return &StateManager{
		stateFile: "terraform.tfstate",
		State:     make(map[string]bool),
	}
}

func (s *StateManager) LoadState() {
	fmt.Println("Loading state...")
	data, err := ioutil.ReadFile(s.stateFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("State file not found, starting with empty state")
			return
		}
		fmt.Printf("Error reading state file: %v\n", err)
		return
	}
	err = json.Unmarshal(data, &s.State)
	if err != nil {
		fmt.Printf("Error parsing state file: %v\n", err)
	}
}

func (s *StateManager) SaveState() {
	fmt.Println("Saving state...")
	data, err := json.Marshal(s.State)
	if err != nil {
		fmt.Printf("Error marshaling state: %v\n", err)
		return
	}
	err = ioutil.WriteFile(s.stateFile, data, 0644)
	if err != nil {
		fmt.Printf("Error writing state file: %v\n", err)
	}
}

func (s *StateManager) ResourceExists(resourceName string) bool {
	return s.State[resourceName]
}

func (s *StateManager) AddResource(resourceName string) {
	s.State[resourceName] = true
}
