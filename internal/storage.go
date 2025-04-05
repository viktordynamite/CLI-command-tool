// Defines data model and storage operations

package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Command represents a stored command line command
type Command struct {
	ID          int       `json:"id"`
	Command     string    `json:"command"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// CommandStore manages the storage of commands
type CommandStore struct {
	Commands []Command `json:"commands"`
	filePath string
}

// NewCommandStore creates a new command store
func NewCommandStore() (*CommandStore, error) {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// creates directory if it doesn't exist
	storeDir := filepath.Join(homeDir, ".cli-command-tool")
	if err := os.MkdirAll(storeDir, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(storeDir, "commands.json")

	store := &CommandStore{
		filePath: filePath,
	}

	// loads existing commands if file exists
	if _, err := os.Stat(filePath); err == nil {
		if err := store.load(); err != nil {
			return nil, err
		}
	}

	return store, nil
}

// load reads commands from the JSON file
func (s *CommandStore) load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s)
}

// save writes commands to the JSON file
func (s *CommandStore) save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

// AddCommand adds a new command to the store
func (s *CommandStore) AddCommand(command, description string) (Command, error) {
	// generates new ID
	newID := 1
	if len(s.Commands) > 0 {
		newID = s.Commands[len(s.Commands)-1].ID + 1
	}

	cmd := Command{
		ID:          newID,
		Command:     command,
		Description: description,
		CreatedAt:   time.Now(),
	}

	s.Commands = append(s.Commands, cmd)

	if err := s.save(); err != nil {
		return Command{}, err
	}

	return cmd, nil
}

// GetAllCommands returns all stored commands
func (s *CommandStore) GetAllCommands() []Command {
	return s.Commands
}

// SearchCommands searches for commands containing the keyword
func (s *CommandStore) SearchCommands(keyword string) []Command {
	var results []Command

	// TODO: Implement search logic

	return results
}
