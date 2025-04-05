// Searches for the command

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/viktordynamite/CLI-command-tool/internal/storage.go"
)

var searcCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search for commands by keyword",
	Long:  `Search through your saved commands using keywords.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		keyword := args[0]

		// To-Do: implement search logic
		fmt.Printf("Searching for commands containing: %s\n", keyword)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searcCmd)
}
