package mymodule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StateManager struct {
	stateFile string
	State     map[string]ResourceState
}

type ResourceState struct {
	Type         string `json:"type"`
	Name         string `json:"name"`
	Ami          string `json:"ami,omitempty"`
	InstanceType string `json:"instance_type,omitempty"`
}

func NewStateManager(stateFile string) *StateManager {
	return &StateManager{
		stateFile: stateFile,
		State:     make(map[string]ResourceState),
	}
}

func (s *StateManager) LoadState() {
	fmt.Printf("Loading state from %s...\n", s.stateFile)
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
	fmt.Printf("Saving state to %s...\n", s.stateFile)
	data, err := json.MarshalIndent(s.State, "", "  ")
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
	_, exists := s.State[resourceName]
	return exists
}

func (s *StateManager) AddResource(resourceName string, resourceState ResourceState) {
	s.State[resourceName] = resourceState
}
