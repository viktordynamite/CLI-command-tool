// Linked file to storage

package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Command represents a stored command with its description
type Command struct {
	Command     string
	Description string
}

// Commands is the global slice of stored commands
var Commands []Command
var CommandsFile string

func init() {
	// Stores commands in user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	// Creates directory if it doesn't exist
	cmdDir := filepath.Join(homeDir, ".cli-commands")
	if _, err := os.Stat(cmdDir); os.IsNotExist(err) {
		os.Mkdir(cmdDir, 0755)
	}

	CommandsFile = filepath.Join(cmdDir, "commands.json")
}

// SaveCommands saves the commands to the JSON file
func SaveCommands() error {
	data, err := json.MarshalIndent(Commands, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(CommandsFile, data, 0644)
}

// LoadCommands loads the commands from the JSON file
func LoadCommands() error {
	data, err := os.ReadFile(CommandsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &Commands)
}
