// File for copy function

package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/viktordynamite/CLI-command-tool/internal/storage"
)

func init() {
	var copyCmd = &cobra.Command{
		Use:   "copy [index]",
		Short: "Copy a command to clipboard",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			index := 0
			fmt.Sscanf(args[0], "%d", &index)

			if index < 1 || index > len(storage.Commands) {
				return fmt.Errorf("invalid command index: %d (valid range: 1-%d)", index, len(storage.Commands))
			}

			command := storage.Commands[index-1].Command
			if err := clipboard.WriteAll(command); err != nil {
				return err
			}

			fmt.Printf("Copied to clipboard: %s\n", command)
			return nil
		},
	}

	rootCmd.AddCommand(copyCmd)
}
