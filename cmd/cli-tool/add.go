// File for adding commands to CLI

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

func init() {
	var addCmd = &cobra.Command{
		Use:   "add [command] [description]",
		Short: "Add a new command",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			command := args[0]
			description := args[1]

			storage.Commands = append(storage.Commands, storage.Command{
				Command:     command,
				Description: description,
			})

			if err := storage.SaveCommands(); err != nil {
				return err
			}

			fmt.Println("Command added successfully")
			return nil
		},
	}

	rootCmd.AddCommand(addCmd)
}
