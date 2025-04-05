// Lists all commands

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/viktordynamite/CLI-command-tool/internal/storage.go"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved commands",
	Long:  `Display a list of all commands you have saved.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// To-Do: implement storage logic to retrieve and display commands
		fmt.Printf("Listing all the saved commands")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
