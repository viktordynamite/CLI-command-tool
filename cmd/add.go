// Adds new commands

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/viktordynamite/CLI-command-tool/internal/storage.go"
)

var addCmd = &cobra.Command{
	Use:   "add [command] [description]",
	Short: "Add a new command to your collection",
	Long:  `Add a new command to your collection with a description for easy reference later.`,
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		commandStr := args[0]
		description := args[1]

		// To-Do: implement storage logic
		fmt.Printf("Added command: %s\nDescription: %s\n", commandStr, description)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
