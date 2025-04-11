// Lists all CLI saved commands

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

func init() {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List all stored commands",
		Run: func(cmd *cobra.Command, args []string) {
			if len(storage.Commands) == 0 {
				fmt.Println("No commands stored. Add some with 'cli-tool add'")
				return
			}

			fmt.Println("Stored commands:")
			fmt.Println("----------------")
			for i, cmd := range storage.Commands {
				fmt.Printf("%d. %s\n   Description: %s\n\n", i+1, cmd.Command, cmd.Description)
			}
		},
	}

	rootCmd.AddCommand(listCmd)
}
