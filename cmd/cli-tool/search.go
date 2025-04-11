// File for search function

package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

func init() {
	var searchCmd = &cobra.Command{
		Use:   "search [keyword]",
		Short: "Search for commands by keyword",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			keyword := args[0]
			found := false
			fmt.Printf("Commands matching '%s':\n", keyword)
			fmt.Println("------------------------")

			for i, cmd := range storage.Commands {
				if strings.Contains(strings.ToLower(cmd.Command), strings.ToLower(keyword)) ||
					strings.Contains(strings.ToLower(cmd.Description), strings.ToLower(keyword)) {
					fmt.Printf("%d. %s\n   Description: %s\n\n", i+1, cmd.Command, cmd.Description)
					found = true
				}
			}

			if !found {
				fmt.Println("No matching commands found.")
			}
		},
	}

	rootCmd.AddCommand(searchCmd)
}
