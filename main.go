// main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

type Command struct {
	Command     string
	Description string
}

var commands []Command
var commandsFile string

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

	commandsFile = filepath.Join(cmdDir, "commands.json")
}

func saveCommands() error {
	// Saves command
	data, err := json.MarshalIndent(commands, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(commandsFile, data, 0644)
}

func loadCommands() error {
	// Loads command
	data, err := os.ReadFile(commandsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &commands)
}

func addCommand(cmd, desc string) {
	// Adds command to CLI
	commands = append(commands, Command{Command: cmd, Description: desc})
	saveCommands()
}

func listCommands() {
	// Lists commands
	if len(commands) == 0 {
		fmt.Println("No commands stored. Add some with 'cli-tool add'")
		return
	}

	fmt.Println("Stored commands:")
	fmt.Println("----------------")
	for i, cmd := range commands {
		fmt.Printf("%d. %s\n   Description: %s\n\n", i+1, cmd.Command, cmd.Description)
	}
}

func searchCommands(keyword string) {
	// Searches for commands by using keyword
	found := false
	fmt.Printf("Commands matching '%s':\n", keyword)
	fmt.Println("------------------------")

	for i, cmd := range commands {
		if strings.Contains(strings.ToLower(cmd.Command), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(cmd.Description), strings.ToLower(keyword)) {
			fmt.Printf("%d. %s\n   Description: %s\n\n", i+1, cmd.Command, cmd.Description)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching commands found.")
	}
}

func copyCommand(index int) error {
	// Copies command
	if index < 1 || index > len(commands) {
		return fmt.Errorf("invalid command index: %d (valid range: 1-%d)", index, len(commands))
	}

	cmd := commands[index-1].Command
	err := clipboard.WriteAll(cmd)
	if err != nil {
		return err
	}

	fmt.Printf("Copied to clipboard: %s\n", cmd)
	return nil
}

func deleteCommand(index int) error {
	// Deletes command
	if index < 1 || index > len(commands) {
		return fmt.Errorf("invalid command index: %d (valid range: 1-%d)", index, len(commands))
	}

	// Removes the command at the specified index
	commands = append(commands[:index-1], commands[index:]...)
	return saveCommands()
}

func executeCommand(index int) error {
	if index < 1 || index > len(commands) {
		return fmt.Errorf("invalid command index: %d (valid range: 1-%d)", index, len(commands))
	}

	cmd := commands[index-1].Command
	fmt.Printf("Executing: %s\n", cmd)

	// On Windows, use cmd.exe or pwsh.exe
	// On Linux use sh

	// More custom stuff later...

	execCmd := exec.Command("cmd", "/C", cmd)
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	execCmd.Stdin = os.Stdin

	return execCmd.Run()
}

func main() {
	if err := loadCommands(); err != nil {
		fmt.Println("Error loading commands:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-tool <command> [args...]")
		fmt.Println("Available commands: add, list, search, copy, delete, exec")
		return
	}

	switch os.Args[1] {
	case "add": // Add
		if len(os.Args) < 4 {
			fmt.Println("Usage: cli-tool add <command> <description>")
			return
		}
		cmd := os.Args[2]
		desc := os.Args[3]
		addCommand(cmd, desc)
		fmt.Println("Command added successfully")

	case "list": // List
		listCommands()

	case "search": // Search
		if len(os.Args) != 3 {
			fmt.Println("Usage: cli-tool search <keyword>")
			return
		}
		searchCommands(os.Args[2])

	case "copy": // Copy
		if len(os.Args) != 3 {
			fmt.Println("Usage: cli-tool copy <index>")
			return
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: index must be a number")
			return
		}

		if err := copyCommand(index); err != nil {
			fmt.Println("Error copying command:", err)
		}

	case "delete": // Delete
		if len(os.Args) != 3 {
			fmt.Println("Usage: cli-tool delete <index>")
			return
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: index must be a number")
			return
		}

		if err := deleteCommand(index); err != nil {
			fmt.Println("Error deleting command:", err)
		} else {
			fmt.Println("Command deleted successfully")
		}

	case "exec": // Execute
		if len(os.Args) != 3 {
			fmt.Println("Usage: cli-tool exec <index>")
			return
		}

		index, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: index must be a number")
			return
		}

		if err := executeCommand(index); err != nil {
			fmt.Println("Error executing command:", err)
		}

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		fmt.Println("Available commands: add, list, search, copy, delete, exec")
	}
}
