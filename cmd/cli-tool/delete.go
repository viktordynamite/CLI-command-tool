// File for delete command

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

func init() {
	var deleteCmd = &cobra.Command{
		Use:   "delete [index]",
		Short: "Delete a command",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := 0
			fmt.Sscanf(args[0], "%d", &index)

			if index < 1 || index > len(storage.Commands) {
				return fmt.Errorf("invalid command index: %d (valid range: 1-%d)", index, len(storage.Commands))
			}

			// Remove the command at the specified index
			storage.Commands = append(storage.Commands[:index-1], storage.Commands[index:]...)

			if err := storage.SaveCommands(); err != nil {
				return err
			}

			fmt.Println("Command deleted successfully")
			return nil
		},
	}

	rootCmd.AddCommand(deleteCmd)
}
